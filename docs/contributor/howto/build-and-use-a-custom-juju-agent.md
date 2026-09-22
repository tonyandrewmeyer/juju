---
myst:
  html_meta:
    description: "Build the Juju agent from source and get it onto a controller, with --build-agent on machine clouds and the operator-update targets on Kubernetes."
---

(build-and-use-a-custom-juju-agent)=
# Build and use a custom Juju agent
> See also: {ref}`compile-and-run-juju-agents-on-different-architectures`

Juju's documentation covers building the client from source. This guide covers
the other half: getting the agent you just built onto a controller, which is
what you need when you're testing a change to Juju itself.

There are two routes and they aren't interchangeable. On machine clouds,
`juju bootstrap --build-agent` compiles the agent out of your source tree and
ships it to the new controller. On Kubernetes that flag is refused, so instead
we build an operator image and import it into the cluster's own image store.

## Before you start

You need a Linux machine you can afford to break, and the two routes want quite
different machines:

* **Machine clouds.** 4 CPUs, 8 GB of memory and 40 GB of disk is enough.
* **Kubernetes.** 8 CPUs, 12 GiB of memory, and more than 40 GB of disk. The
  Kubernetes route needs the build and a running cluster on the same machine,
  so the machine-cloud numbers don't carry over.

> See more: {ref}`custom-agent-sizing`

Install `build-essential` first:

```text
sudo apt install build-essential
```

<!-- REVIEW: the two install-dependencies bugs. FINDINGS says these are to be
     raised with the Juju team as bugs, and that the how-to must not quietly
     carry a workaround as though it were the normal path. The block below
     names them as bugs and gives the error output, rather than folding the
     workarounds into the steps. Once the bugs are filed, link them from here;
     once they're fixed, this block comes out. -->

```{caution}
Two things in `make install-dependencies` are broken, on both `main` and `3.6`,
and both of them stop you on a clean machine. They are bugs in the Makefile
rather than something you've done wrong.

1. **The Go snap channel is malformed.** The target installs the Go snap from
   `$(GO_MOD_VERSION)/stable`, and `GO_MOD_VERSION` is the raw value out of
   `go.mod`, currently `1.26.6`. There is no `1.26.6/stable` channel, because
   the snap tracks are `1.26/stable`, so on a machine without Go already
   installed the first documented step fails:

   ```text
   error: snap "go" is not available on 1.26.6/stable but is available to install ...
   make: *** [Makefile:595: install-snap-dependencies] Error 1
   ```

   Until that's fixed, install Go yourself first:
   `sudo snap install go --channel=1.26/stable --classic`.

2. **`build-essential` isn't installed**, which is why it's a manual step
   above. Without it you don't have `make` to run the target with in the first
   place, and `make install` later fails at `rebuild-schema` with
   `cgo: C compiler "gcc" not found`.
```

## Build and install the binaries

```text
make install-dependencies
make install
```

`make install` puts the binaries in `$GOBIN`, which is `~/go/bin` unless you've
set it to something else. Put that directory early on your `PATH`, so that the
`juju` you run is the one you just built:

```text
export PATH=~/go/bin:$PATH
juju version
```

On `main` you get `juju`, `jujuc`, `jujuagentd`, `containeragent`,
`juju-metadata` and `pebble`. On `3.6` the agent binary is `jujud` rather than
`jujuagentd`.

Expect this to take a while. On `main`, `make install` took 18 minutes on 4
CPUs with a warm Go module cache. Not all of that is compilation: building
`jujuagentd` on `main` downloads a musl tarball of about 510 MB, and the dqlite
dependencies after it. On a slow connection that download is most of the wait.

```{note}
If you're going to build both `main` and `3.6` on one machine, give each of
them its own `GOBIN` and its own `JUJU_DATA`. The `*-operator-update` targets
depend on `host-install`, which reinstalls the client into `$GOBIN`, so without
that separation the second build replaces the first one's client.
```

## Machine clouds

### Bootstrap a new controller

Run `juju bootstrap` from inside the Juju source tree, with `--build-agent`:

```text
cd ~/juju
juju bootstrap localhost lxdtest --build-agent
```

The `cd` isn't optional. Juju builds the agent from the tree you're standing
in, and refuses to build one from anywhere else:

```text
ERROR failed to bootstrap model: cannot package bootstrap agent binary: cannot
build jujuagentd agent binary from source: cannot build juju agent outside of
github.com/juju/juju tree
```

> See also: {ref}`command-juju-bootstrap`

### Upgrade a controller you already have

<!-- REVIEW: unverified. FINDINGS lists "upgrade-controller --build-agent" under
     "Still not verified" - the flag and its help text were read from the source
     (cmd/juju/commands/upgradecontroller.go), but the command was never run. -->

On the second and later iterations you usually don't want a fresh controller.
`juju upgrade-controller --build-agent` builds the agent the same way and puts
it onto a controller you already have:

```text
cd ~/juju
juju upgrade-controller --build-agent
```

The flag's own help text says "for development use only", which is exactly what
this is.

> See also: {ref}`command-juju-upgrade-controller`

If the controller isn't the same architecture or operating system as your host,
`--build-agent` won't help you, and you want a local simplestreams repository
instead. That guide already exists:
{ref}`compile-and-run-juju-agents-on-different-architectures`.

## Kubernetes

`--build-agent` is not an option here. Bootstrapping a Kubernetes controller
with it fails with `--build-agent when bootstrapping a k8s controller not
supported`.

<!-- REVIEW: not in FINDINGS. Read from cmd/juju/commands/upgradecontroller.go
     (NotSupportedf("--build-agent for k8s model upgrades")) on this session's
     checkout of main, b4892e498a. Not run. Drop the sentence if you'd rather
     the draft cite only what was executed. -->

The upgrade route is closed off too: `upgrade-controller --build-agent` returns
`--build-agent for k8s model upgrades not supported`.

What replaces both is building the operator image and importing it into the
image store that your cluster reads. The Makefile has a target per cluster
flavour:

| Cluster | Target |
|---|---|
| Canonical K8s | `ck8s-operator-update` |
| MicroK8s | `microk8s-operator-update` |
| k3s | `k3s-operator-update` |
| minikube | `minikube-operator-update` |
| A Charmed Kubernetes model (set `JUJU_K8S_MODEL`) | `local-operator-update` |

<!-- REVIEW: only ck8s-operator-update was run (FINDINGS, "Verified on Canonical
     K8s" and "Verified on 3.6"). The microk8s, k3s, minikube and local targets
     are listed from the Makefile and are under "Still not verified"; the
     microk8s one in particular streams via microk8s.ctr and has a separate
     macOS branch in make_functions.sh, so it isn't the same command shape.
     Everything below this table is Canonical K8s. -->

The rest of this section uses Canonical K8s.

### Use podman, not Docker

<!-- REVIEW: the brief's outline says "Docker required" for this branch. That
     follows the old ops snippet rather than FINDINGS, which found the opposite:
     the Makefile prefers podman, and Docker actively breaks the Canonical K8s
     bootstrap. Written to match FINDINGS. -->

The `*-operator-update` targets build the image with `$(OCI_BUILDER)`, which is
`podman` if you have it and `docker` if you don't. Use podman.

Docker and Canonical K8s both want `/run/containerd`. With `docker.io`
installed, `k8s bootstrap` refuses to run at all:

```text
The path '/run/containerd' required for the containerd socket already exists.
This may mean that another service is already using that path, and it conflicts
with the k8s snap.
```

The `containerd-base-dir` option that the error points you towards is worse
than the problem it solves. `ck8s-operator-update` hardcodes
`sudo /snap/k8s/current/bin/ctr -n k8s.io`, and `ctr` talks to
`/run/containerd/containerd.sock`. Move the cluster's containerd somewhere else
and your image is imported into Docker's containerd instead. Nothing fails, and
the controller quietly comes up running a stock agent.

Podman doesn't take `/run/containerd`, so none of this arises.

### Add the cluster as a cloud

Canonical K8s isn't detected automatically, the way MicroK8s and LXD are, so
`juju bootstrap k8s` gives you `ERROR unknown cloud "k8s"`. Register it
yourself:

```text
sudo k8s config > ~/.kube/config
juju add-k8s ck8s --client
```

> See also: {ref}`command-juju-add-k8s`

### Build the image and bootstrap

```text
JUJU_BUILD_NUMBER=1 make ck8s-operator-update
juju bootstrap ck8s ctest
```

On `main` the target took 2 minutes and the bootstrap 43 seconds. The target
ends by piping the image into the cluster's containerd:

```text
podman save "ghcr.io/juju/jujud-operator:4.2-beta1.1" | sudo /snap/k8s/current/bin/ctr -n k8s.io images import -
unpacking ghcr.io/juju/jujud-operator:4.2-beta1.1 (sha256:73cb708e...)...done
```

`JUJU_BUILD_NUMBER` is what puts the trailing `.1` on that tag. Strictly you
don't need it when you build from a branch tip, because a branch tip always
carries the next, unreleased patch version, so there's no published image to
collide with: `4.2-beta1` and `3.6.29` are both 404 on ghcr.io while `3.6.28`
is there. Set it anyway, for two reasons. If you build a release tag such as
`v3.6.28` rather than a branch tip, a stray registry pull would be
indistinguishable from your own image. And the build number is the cheapest
confirmation you have that you're running your own agent, because it shows up
in `juju controllers` as well as in the tag.

Only the operator image comes from your build. The controller pod's `charm`
container runs `ghcr.io/juju/charm-base:ubuntu-24.04`, which is fetched from
the registry, so the bootstrap still needs network.

(custom-agent-sizing)=
### Sizing

8 GiB of memory is not enough. Linking `jujuagentd` while the cluster was
running drove an 8 GiB machine to a load average of 95 and exhausted its
memory; SSH became unreachable and the machine had to be stopped. With 12 GiB,
and the cluster stopped for the build, the build took about 3 minutes.

40 GB of disk is not enough either. The controller's PVC asks for 20Gi, and the
Go build cache, the source tree and `~/go` between them can easily leave less
than that free. When the claim can't be satisfied, the cluster says:

```text
ProvisioningFailed  failed to provision volume with StorageClass "csi-rawfile-default":
  rpc error: code = ResourceExhausted desc = Not enough disk space
```

Juju doesn't show you that. What you see is `timed out waiting for controller
pod: client rate limiter Wait returned an error`, 20 minutes later, with the
namespace torn down behind it. Bootstrap with `--keep-broken` to keep the
namespace and read the real error. `go clean -cache` is the quickest way to
free space, and it freed 26 GiB here.

One more trap: the rawfile CSI reserves a PVC's logical size rather than what's
actually in it. A second bootstrap failed with `ResourceExhausted` at 24 GiB
free, because the first controller's 20Gi claim (holding 391M of real data)
left too little for another one. Two controllers on one node want 40Gi or more
of backing disk, whatever thin provisioning suggests.

## Check that the controller is running your build

On any cloud, `juju controllers` reports the agent version, and your build has
the build number on the end:

```text
agent-version: 4.2-beta1.1
```

A released agent has no trailing `.1`, so that digit is the whole check.

<!-- REVIEW: FINDINGS records the kubelet event below, and that it was read off
     the controller pod, but not the command used to read it. The namespace is
     the usual controller-<controller-name>. Check the command before this is
     published. -->

On Kubernetes there's a stronger check available. Ask the controller pod which
image it actually used:

```text
kubectl -n controller-ctest describe pod controller-0
```

```text
Normal  Pulled  pod/controller-0
  Container image "ghcr.io/juju/jujud-operator:4.2-beta1.1" already present on machine
```

"already present on machine" is the line that matters: the kubelet found the
image you imported and didn't go to the registry. `imagePullPolicy` is
`IfNotPresent`. The `api-server` container and both init containers
(`controller-config-seed` and `charm-init`) run that image.

## Juju 3.6

<!-- REVIEW: "all of the above" overreaches slightly. FINDINGS ran only the
     Kubernetes leg on 3.6; the machine-cloud steps (--build-agent against LXD)
     were verified on main only. Nothing in the source suggests they differ
     beyond the jujud name, but nobody has run them. -->

All of the above works on `3.6`, with these differences:

* The agent binary is `jujud`, not `jujuagentd`. On `main` the operator image
  adds a `ln -s jujuagentd /opt/jujud` compatibility symlink; on `3.6` there's
  nothing to bridge.
* `install-dependencies` installs the `juju-db` snap (4.4.30) and a list of apt
  packages, where `main` installs only `libsqlite3-dev`. There's no musl or
  dqlite download, so the 510 MB above doesn't apply.
* The Go snap channel bug is still there, but you may not see it. If you've
  already built `main` on the same machine then Go is installed, and the target
  reports `Using installed go-1.26.6` instead of failing.
* Everything is slower and bigger. `make install` took 14 minutes from a cold
  cache, against 3 minutes warm on `main`, and `ck8s-operator-update` took 9
  minutes and produced a 553.2 MiB image against 454.1 MiB. The `3.6`
  Dockerfile is a two-stage, 22-step build; `main`'s is 8 steps.
* The controller pod carries Mongo: one init container (`charm-init`) and three
  containers (`charm`, `mongodb`, `api-server`), against two init containers
  (`controller-config-seed`, `charm-init`) and two containers on `main`. The
  `mongodb` container runs `ghcr.io/juju/juju-db:4.4.30`, pulled from the
  registry, so the `3.6` bootstrap needs network for that as well as for
  `charm-base`.
