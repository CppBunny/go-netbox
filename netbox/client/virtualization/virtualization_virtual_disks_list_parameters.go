// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewVirtualizationVirtualDisksListParams creates a new VirtualizationVirtualDisksListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualDisksListParams() *VirtualizationVirtualDisksListParams {
	return &VirtualizationVirtualDisksListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualDisksListParamsWithTimeout creates a new VirtualizationVirtualDisksListParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualDisksListParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualDisksListParams {
	return &VirtualizationVirtualDisksListParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualDisksListParamsWithContext creates a new VirtualizationVirtualDisksListParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualDisksListParamsWithContext(ctx context.Context) *VirtualizationVirtualDisksListParams {
	return &VirtualizationVirtualDisksListParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualDisksListParamsWithHTTPClient creates a new VirtualizationVirtualDisksListParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualDisksListParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualDisksListParams {
	return &VirtualizationVirtualDisksListParams{
		HTTPClient: client,
	}
}

/*
VirtualizationVirtualDisksListParams contains all the parameters to send to the API endpoint

	for the virtualization virtual disks list operation.

	Typically these are written to a http.Request.
*/
type VirtualizationVirtualDisksListParams struct {

	// Created.
	Created *string

	// CreatedGt.
	CreatedGt *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLt.
	CreatedLt *string

	// CreatedLte.
	CreatedLte *string

	// Createdn.
	Createdn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGt.
	LastUpdatedGt *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLt.
	LastUpdatedLt *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	// LastUpdatedn.
	LastUpdatedn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Master.
	Master *string

	// MasterID.
	MasterID *string

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	/* Ordering.

	   Which field to use when ordering the results.
	*/
	Ordering *string

	// Q.
	Q *string

	// Region.
	Region *string

	// RegionID.
	RegionID *string

	// Site.
	Site *string

	// SiteGroup.
	SiteGroup *string

	// SiteGroupID.
	SiteGroupID *string

	// SiteID.
	SiteID *string

	// Tag.
	Tag []string

	// Tagn.
	Tagn *string

	// Tenant.
	Tenant *string

	// TenantID.
	TenantID *string

	// VirtualMachinen.
	VirtualMachinen *string

	// VirtualMachineIDn.
	VirtualMachineIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual disks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualDisksListParams) WithDefaults() *VirtualizationVirtualDisksListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual disks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualDisksListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualDisksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithContext(ctx context.Context) *VirtualizationVirtualDisksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualDisksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithCreated(created *string) *VirtualizationVirtualDisksListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGt adds the createdGt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithCreatedGt(createdGt *string) *VirtualizationVirtualDisksListParams {
	o.SetCreatedGt(createdGt)
	return o
}

// SetCreatedGt adds the createdGt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetCreatedGt(createdGt *string) {
	o.CreatedGt = createdGt
}

// WithCreatedGte adds the createdGte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithCreatedGte(createdGte *string) *VirtualizationVirtualDisksListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLt adds the createdLt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithCreatedLt(createdLt *string) *VirtualizationVirtualDisksListParams {
	o.SetCreatedLt(createdLt)
	return o
}

// SetCreatedLt adds the createdLt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetCreatedLt(createdLt *string) {
	o.CreatedLt = createdLt
}

// WithCreatedLte adds the createdLte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithCreatedLte(createdLte *string) *VirtualizationVirtualDisksListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithCreatedn adds the createdn to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithCreatedn(createdn *string) *VirtualizationVirtualDisksListParams {
	o.SetCreatedn(createdn)
	return o
}

// SetCreatedn adds the createdN to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetCreatedn(createdn *string) {
	o.Createdn = createdn
}

// WithID adds the id to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithID(id *string) *VirtualizationVirtualDisksListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithIDGt(iDGt *string) *VirtualizationVirtualDisksListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithIDGte(iDGte *string) *VirtualizationVirtualDisksListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithIDLt(iDLt *string) *VirtualizationVirtualDisksListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithIDLte(iDLte *string) *VirtualizationVirtualDisksListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithIDn(iDn *string) *VirtualizationVirtualDisksListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithLastUpdated(lastUpdated *string) *VirtualizationVirtualDisksListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGt adds the lastUpdatedGt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithLastUpdatedGt(lastUpdatedGt *string) *VirtualizationVirtualDisksListParams {
	o.SetLastUpdatedGt(lastUpdatedGt)
	return o
}

// SetLastUpdatedGt adds the lastUpdatedGt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetLastUpdatedGt(lastUpdatedGt *string) {
	o.LastUpdatedGt = lastUpdatedGt
}

// WithLastUpdatedGte adds the lastUpdatedGte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithLastUpdatedGte(lastUpdatedGte *string) *VirtualizationVirtualDisksListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLt adds the lastUpdatedLt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithLastUpdatedLt(lastUpdatedLt *string) *VirtualizationVirtualDisksListParams {
	o.SetLastUpdatedLt(lastUpdatedLt)
	return o
}

// SetLastUpdatedLt adds the lastUpdatedLt to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetLastUpdatedLt(lastUpdatedLt *string) {
	o.LastUpdatedLt = lastUpdatedLt
}

// WithLastUpdatedLte adds the lastUpdatedLte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithLastUpdatedLte(lastUpdatedLte *string) *VirtualizationVirtualDisksListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLastUpdatedn adds the lastUpdatedn to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithLastUpdatedn(lastUpdatedn *string) *VirtualizationVirtualDisksListParams {
	o.SetLastUpdatedn(lastUpdatedn)
	return o
}

// SetLastUpdatedn adds the lastUpdatedN to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetLastUpdatedn(lastUpdatedn *string) {
	o.LastUpdatedn = lastUpdatedn
}

// WithLimit adds the limit to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithLimit(limit *int64) *VirtualizationVirtualDisksListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaster adds the master to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithMaster(master *string) *VirtualizationVirtualDisksListParams {
	o.SetMaster(master)
	return o
}

// SetMaster adds the master to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetMaster(master *string) {
	o.Master = master
}

// WithMasterID adds the masterID to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithMasterID(masterID *string) *VirtualizationVirtualDisksListParams {
	o.SetMasterID(masterID)
	return o
}

// SetMasterID adds the masterId to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetMasterID(masterID *string) {
	o.MasterID = masterID
}

// WithName adds the name to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithName(name *string) *VirtualizationVirtualDisksListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameEmpty(nameEmpty *string) *VirtualizationVirtualDisksListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameIc(nameIc *string) *VirtualizationVirtualDisksListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameIe(nameIe *string) *VirtualizationVirtualDisksListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameIew(nameIew *string) *VirtualizationVirtualDisksListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameIsw(nameIsw *string) *VirtualizationVirtualDisksListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNamen(namen *string) *VirtualizationVirtualDisksListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameNic(nameNic *string) *VirtualizationVirtualDisksListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameNie(nameNie *string) *VirtualizationVirtualDisksListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameNiew(nameNiew *string) *VirtualizationVirtualDisksListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithNameNisw(nameNisw *string) *VirtualizationVirtualDisksListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithOffset(offset *int64) *VirtualizationVirtualDisksListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithOrdering(ordering *string) *VirtualizationVirtualDisksListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithQ adds the q to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithQ(q *string) *VirtualizationVirtualDisksListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithRegion(region *string) *VirtualizationVirtualDisksListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionID adds the regionID to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithRegionID(regionID *string) *VirtualizationVirtualDisksListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithSite adds the site to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithSite(site *string) *VirtualizationVirtualDisksListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteGroup adds the siteGroup to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithSiteGroup(siteGroup *string) *VirtualizationVirtualDisksListParams {
	o.SetSiteGroup(siteGroup)
	return o
}

// SetSiteGroup adds the siteGroup to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetSiteGroup(siteGroup *string) {
	o.SiteGroup = siteGroup
}

// WithSiteGroupID adds the siteGroupID to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithSiteGroupID(siteGroupID *string) *VirtualizationVirtualDisksListParams {
	o.SetSiteGroupID(siteGroupID)
	return o
}

// SetSiteGroupID adds the siteGroupId to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetSiteGroupID(siteGroupID *string) {
	o.SiteGroupID = siteGroupID
}

// WithSiteID adds the siteID to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithSiteID(siteID *string) *VirtualizationVirtualDisksListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithTag adds the tag to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithTag(tag []string) *VirtualizationVirtualDisksListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetTag(tag []string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithTagn(tagn *string) *VirtualizationVirtualDisksListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithTenant(tenant *string) *VirtualizationVirtualDisksListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithTenantID(tenantID *string) *VirtualizationVirtualDisksListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithVirtualMachinen adds the virtualMachinen to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithVirtualMachinen(virtualMachinen *string) *VirtualizationVirtualDisksListParams {
	o.SetVirtualMachinen(virtualMachinen)
	return o
}

// SetVirtualMachinen adds the virtualMachineN to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetVirtualMachinen(virtualMachinen *string) {
	o.VirtualMachinen = virtualMachinen
}

// WithVirtualMachineIDn adds the virtualMachineIDn to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) WithVirtualMachineIDn(virtualMachineIDn *string) *VirtualizationVirtualDisksListParams {
	o.SetVirtualMachineIDn(virtualMachineIDn)
	return o
}

// SetVirtualMachineIDn adds the virtualMachineIdN to the virtualization virtual disks list params
func (o *VirtualizationVirtualDisksListParams) SetVirtualMachineIDn(virtualMachineIDn *string) {
	o.VirtualMachineIDn = virtualMachineIDn
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualDisksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGt != nil {

		// query param created__gt
		var qrCreatedGt string

		if o.CreatedGt != nil {
			qrCreatedGt = *o.CreatedGt
		}
		qCreatedGt := qrCreatedGt
		if qCreatedGt != "" {

			if err := r.SetQueryParam("created__gt", qCreatedGt); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLt != nil {

		// query param created__lt
		var qrCreatedLt string

		if o.CreatedLt != nil {
			qrCreatedLt = *o.CreatedLt
		}
		qCreatedLt := qrCreatedLt
		if qCreatedLt != "" {

			if err := r.SetQueryParam("created__lt", qCreatedLt); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Createdn != nil {

		// query param created__n
		var qrCreatedn string

		if o.Createdn != nil {
			qrCreatedn = *o.Createdn
		}
		qCreatedn := qrCreatedn
		if qCreatedn != "" {

			if err := r.SetQueryParam("created__n", qCreatedn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGt != nil {

		// query param last_updated__gt
		var qrLastUpdatedGt string

		if o.LastUpdatedGt != nil {
			qrLastUpdatedGt = *o.LastUpdatedGt
		}
		qLastUpdatedGt := qrLastUpdatedGt
		if qLastUpdatedGt != "" {

			if err := r.SetQueryParam("last_updated__gt", qLastUpdatedGt); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLt != nil {

		// query param last_updated__lt
		var qrLastUpdatedLt string

		if o.LastUpdatedLt != nil {
			qrLastUpdatedLt = *o.LastUpdatedLt
		}
		qLastUpdatedLt := qrLastUpdatedLt
		if qLastUpdatedLt != "" {

			if err := r.SetQueryParam("last_updated__lt", qLastUpdatedLt); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedn != nil {

		// query param last_updated__n
		var qrLastUpdatedn string

		if o.LastUpdatedn != nil {
			qrLastUpdatedn = *o.LastUpdatedn
		}
		qLastUpdatedn := qrLastUpdatedn
		if qLastUpdatedn != "" {

			if err := r.SetQueryParam("last_updated__n", qLastUpdatedn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Master != nil {

		// query param master
		var qrMaster string

		if o.Master != nil {
			qrMaster = *o.Master
		}
		qMaster := qrMaster
		if qMaster != "" {

			if err := r.SetQueryParam("master", qMaster); err != nil {
				return err
			}
		}
	}

	if o.MasterID != nil {

		// query param master_id
		var qrMasterID string

		if o.MasterID != nil {
			qrMasterID = *o.MasterID
		}
		qMasterID := qrMasterID
		if qMasterID != "" {

			if err := r.SetQueryParam("master_id", qMasterID); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string

		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {

			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Region != nil {

		// query param region
		var qrRegion string

		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {

			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}
	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string

		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {

			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}
	}

	if o.Site != nil {

		// query param site
		var qrSite string

		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {

			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}
	}

	if o.SiteGroup != nil {

		// query param site_group
		var qrSiteGroup string

		if o.SiteGroup != nil {
			qrSiteGroup = *o.SiteGroup
		}
		qSiteGroup := qrSiteGroup
		if qSiteGroup != "" {

			if err := r.SetQueryParam("site_group", qSiteGroup); err != nil {
				return err
			}
		}
	}

	if o.SiteGroupID != nil {

		// query param site_group_id
		var qrSiteGroupID string

		if o.SiteGroupID != nil {
			qrSiteGroupID = *o.SiteGroupID
		}
		qSiteGroupID := qrSiteGroupID
		if qSiteGroupID != "" {

			if err := r.SetQueryParam("site_group_id", qSiteGroupID); err != nil {
				return err
			}
		}
	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string

		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {

			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// binding items for tag
		joinedTag := o.bindParamTag(reg)

		// query array param tag
		if err := r.SetQueryParam("tag", joinedTag...); err != nil {
			return err
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachinen != nil {

		// query param virtual_machine__n
		var qrVirtualMachinen string

		if o.VirtualMachinen != nil {
			qrVirtualMachinen = *o.VirtualMachinen
		}
		qVirtualMachinen := qrVirtualMachinen
		if qVirtualMachinen != "" {

			if err := r.SetQueryParam("virtual_machine__n", qVirtualMachinen); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineIDn != nil {

		// query param virtual_machine_id__n
		var qrVirtualMachineIDn string

		if o.VirtualMachineIDn != nil {
			qrVirtualMachineIDn = *o.VirtualMachineIDn
		}
		qVirtualMachineIDn := qrVirtualMachineIDn
		if qVirtualMachineIDn != "" {

			if err := r.SetQueryParam("virtual_machine_id__n", qVirtualMachineIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVirtualizationVirtualDisksList binds the parameter tag
func (o *VirtualizationVirtualDisksListParams) bindParamTag(formats strfmt.Registry) []string {
	tagIR := o.Tag

	var tagIC []string
	for _, tagIIR := range tagIR { // explode []string

		tagIIV := tagIIR // string as string
		tagIC = append(tagIC, tagIIV)
	}

	// items.CollectionFormat: "multi"
	tagIS := swag.JoinByFormat(tagIC, "multi")

	return tagIS
}