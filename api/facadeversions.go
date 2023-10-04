// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import "github.com/juju/collections/set"

// facadeVersions lists the best version of facades that we want to support. This
// will be used to pick out a default version for communication, given the list
// of known versions that the API server tells us it is capable of supporting.
// This map should be updated whenever the API server exposes a new version (so
// that the client will use it whenever it is available). Removal of an API
// server version can be done when any prior releases are no longer supported.
// Normally, this can be done at major release, although additional thought around
// FullStatus (client facade) and Migration (controller facade) is needed.
// New facades should start at 1.
// We no longer support facade versions at 0.
var facadeVersions = map[string][]int{
	"Action":                       {7},
	"ActionPruner":                 {1},
	"Agent":                        {3},
	"AgentLifeFlag":                {1},
	"AgentTools":                   {1},
	"AllModelWatcher":              {4},
	"AllWatcher":                   {3},
	"Annotations":                  {2},
	"Application":                  {15, 16, 17, 18, 19},
	"ApplicationOffers":            {4},
	"ApplicationScaler":            {1},
	"Backups":                      {3},
	"Block":                        {2},
	"Bundle":                       {6},
	"CAASAgent":                    {2},
	"CAASAdmission":                {1},
	"CAASApplication":              {1},
	"CAASApplicationProvisioner":   {1},
	"CAASModelConfigManager":       {1},
	"CAASFirewaller":               {1},
	"CAASFirewallerSidecar":        {1},
	"CAASModelOperator":            {1},
	"CAASOperator":                 {1},
	"CAASOperatorProvisioner":      {1},
	"CAASOperatorUpgrader":         {1},
	"CAASUnitProvisioner":          {2},
	"CharmDownloader":              {1},
	"CharmRevisionUpdater":         {2},
	"Charms":                       {5, 6, 7},
	"Cleaner":                      {2},
	"Client":                       {6, 7},
	"Cloud":                        {7},
	"Controller":                   {11},
	"CredentialManager":            {1},
	"CredentialValidator":          {2},
	"CrossController":              {1},
	"CrossModelRelations":          {2},
	"CrossModelSecrets":            {1},
	"Deployer":                     {1},
	"DiskManager":                  {2},
	"EntityWatcher":                {2},
	"EnvironUpgrader":              {1},
	"ExternalControllerUpdater":    {1},
	"FanConfigurer":                {1},
	"FilesystemAttachmentsWatcher": {2},
	"Firewaller":                   {7},
	"HighAvailability":             {2},
	"HostKeyReporter":              {1},
	"ImageMetadata":                {3},
	"ImageMetadataManager":         {1},
	"InstanceMutater":              {3},
	"InstancePoller":               {4},
	"KeyManager":                   {1},
	"KeyUpdater":                   {1},
	"LeadershipService":            {2},
	"LifeFlag":                     {1},
	"LogForwarding":                {1},
	"Logger":                       {1},
	"MachineActions":               {1},
	"MachineManager":               {9, 10},
	"MachineUndertaker":            {1},
	"Machiner":                     {5},
	"MeterStatus":                  {2},
	"MetricsAdder":                 {2},
	"MetricsDebug":                 {2},
	"MetricsManager":               {1},
	"MigrationFlag":                {1},
	"MigrationMaster":              {3},
	"MigrationMinion":              {1},
	"MigrationStatusWatcher":       {1},
	"MigrationTarget":              {1, 2},
	"ModelConfig":                  {3},
	"ModelGeneration":              {4},
	"ModelManager":                 {9},
	"ModelSummaryWatcher":          {1},
	"ModelUpgrader":                {1},
	"NotifyWatcher":                {1},
	"OfferStatusWatcher":           {1},
	"Payloads":                     {1},
	"PayloadsHookContext":          {1},
	"Pinger":                       {1},
	"Provisioner":                  {11},
	"ProxyUpdater":                 {2},
	"Reboot":                       {2},
	"RelationStatusWatcher":        {1},
	"RelationUnitsWatcher":         {1},
	"RemoteRelations":              {2},
	"RemoteRelationWatcher":        {1},
	"Resources":                    {3},
	"ResourcesHookContext":         {1},
	"RetryStrategy":                {1},
	"SecretsTriggerWatcher":        {1},
	"SecretBackends":               {1},
	"SecretBackendsManager":        {1},
	"SecretBackendsRotateWatcher":  {1},
	"SecretsRevisionWatcher":       {1},
	"Secrets":                      {1, 2},
	"SecretsManager":               {1, 2},
	"SecretsDrain":                 {1},
	"Singular":                     {2},
	"Spaces":                       {6},
	"SSHClient":                    {4},
	"StatusHistory":                {2},
	"Storage":                      {6},
	"StorageProvisioner":           {4},
	"StringsWatcher":               {1},
	"Subnets":                      {5},
	"Undertaker":                   {1},
	"UnitAssigner":                 {1},
	"Uniter":                       {18, 19},
	"Upgrader":                     {1},
	"UpgradeSeries":                {3},
	"UpgradeSteps":                 {2},
	"UserManager":                  {3},
	"VolumeAttachmentsWatcher":     {2},
	"VolumeAttachmentPlansWatcher": {1},
}

// bestVersion tries to find the newest version in the version list that we can
// use.
func bestVersion(desired []int, versions []int) int {
	intersection := set.NewInts(desired...).Intersection(set.NewInts(versions...))
	if intersection.Size() == 0 {
		return 0
	}
	sorted := intersection.SortedValues()
	return sorted[len(sorted)-1]
}
