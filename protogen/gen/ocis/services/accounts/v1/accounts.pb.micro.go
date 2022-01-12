// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: ocis/services/accounts/v1/accounts.proto

package v1

import (
	fmt "fmt"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v1 "github.com/owncloud/ocis/protogen/gen/ocis/messages/accounts/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for AccountsService service

func NewAccountsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "AccountsService.ListAccounts",
			Path:    []string{"/api/v0/accounts/accounts-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "AccountsService.GetAccount",
			Path:    []string{"/api/v0/accounts/accounts-get"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "AccountsService.CreateAccount",
			Path:    []string{"/api/v0/accounts/accounts-create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "AccountsService.UpdateAccount",
			Path:    []string{"/api/v0/accounts/accounts-update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "AccountsService.DeleteAccount",
			Path:    []string{"/api/v0/accounts/accounts-delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for AccountsService service

type AccountsService interface {
	// Lists accounts
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...client.CallOption) (*ListAccountsResponse, error)
	// Gets an account
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*v1.Account, error)
	// Creates an account
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...client.CallOption) (*v1.Account, error)
	// Updates an account
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...client.CallOption) (*v1.Account, error)
	// Deletes an account
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type accountsService struct {
	c    client.Client
	name string
}

func NewAccountsService(name string, c client.Client) AccountsService {
	return &accountsService{
		c:    c,
		name: name,
	}
}

func (c *accountsService) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...client.CallOption) (*ListAccountsResponse, error) {
	req := c.c.NewRequest(c.name, "AccountsService.ListAccounts", in)
	out := new(ListAccountsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*v1.Account, error) {
	req := c.c.NewRequest(c.name, "AccountsService.GetAccount", in)
	out := new(v1.Account)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...client.CallOption) (*v1.Account, error) {
	req := c.c.NewRequest(c.name, "AccountsService.CreateAccount", in)
	out := new(v1.Account)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...client.CallOption) (*v1.Account, error) {
	req := c.c.NewRequest(c.name, "AccountsService.UpdateAccount", in)
	out := new(v1.Account)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "AccountsService.DeleteAccount", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountsService service

type AccountsServiceHandler interface {
	// Lists accounts
	ListAccounts(context.Context, *ListAccountsRequest, *ListAccountsResponse) error
	// Gets an account
	GetAccount(context.Context, *GetAccountRequest, *v1.Account) error
	// Creates an account
	CreateAccount(context.Context, *CreateAccountRequest, *v1.Account) error
	// Updates an account
	UpdateAccount(context.Context, *UpdateAccountRequest, *v1.Account) error
	// Deletes an account
	DeleteAccount(context.Context, *DeleteAccountRequest, *emptypb.Empty) error
}

func RegisterAccountsServiceHandler(s server.Server, hdlr AccountsServiceHandler, opts ...server.HandlerOption) error {
	type accountsService interface {
		ListAccounts(ctx context.Context, in *ListAccountsRequest, out *ListAccountsResponse) error
		GetAccount(ctx context.Context, in *GetAccountRequest, out *v1.Account) error
		CreateAccount(ctx context.Context, in *CreateAccountRequest, out *v1.Account) error
		UpdateAccount(ctx context.Context, in *UpdateAccountRequest, out *v1.Account) error
		DeleteAccount(ctx context.Context, in *DeleteAccountRequest, out *emptypb.Empty) error
	}
	type AccountsService struct {
		accountsService
	}
	h := &accountsServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "AccountsService.ListAccounts",
		Path:    []string{"/api/v0/accounts/accounts-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "AccountsService.GetAccount",
		Path:    []string{"/api/v0/accounts/accounts-get"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "AccountsService.CreateAccount",
		Path:    []string{"/api/v0/accounts/accounts-create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "AccountsService.UpdateAccount",
		Path:    []string{"/api/v0/accounts/accounts-update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "AccountsService.DeleteAccount",
		Path:    []string{"/api/v0/accounts/accounts-delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&AccountsService{h}, opts...))
}

type accountsServiceHandler struct {
	AccountsServiceHandler
}

func (h *accountsServiceHandler) ListAccounts(ctx context.Context, in *ListAccountsRequest, out *ListAccountsResponse) error {
	return h.AccountsServiceHandler.ListAccounts(ctx, in, out)
}

func (h *accountsServiceHandler) GetAccount(ctx context.Context, in *GetAccountRequest, out *v1.Account) error {
	return h.AccountsServiceHandler.GetAccount(ctx, in, out)
}

func (h *accountsServiceHandler) CreateAccount(ctx context.Context, in *CreateAccountRequest, out *v1.Account) error {
	return h.AccountsServiceHandler.CreateAccount(ctx, in, out)
}

func (h *accountsServiceHandler) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, out *v1.Account) error {
	return h.AccountsServiceHandler.UpdateAccount(ctx, in, out)
}

func (h *accountsServiceHandler) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, out *emptypb.Empty) error {
	return h.AccountsServiceHandler.DeleteAccount(ctx, in, out)
}

// Api Endpoints for GroupsService service

func NewGroupsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "GroupsService.ListGroups",
			Path:    []string{"/api/v0/accounts/groups-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "GroupsService.GetGroup",
			Path:    []string{"/api/v0/accounts/groups-get"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "GroupsService.CreateGroup",
			Path:    []string{"/api/v0/accounts/groups-create"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "GroupsService.UpdateGroup",
			Path:    []string{"/api/v0/accounts/groups-update"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "GroupsService.DeleteGroup",
			Path:    []string{"/api/v0/accounts/groups-delete"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "GroupsService.AddMember",
			Path:    []string{"/api/v0/groups/{group_id=*}/members/$ref"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "GroupsService.RemoveMember",
			Path:    []string{"/api/v0/groups/{group_id=*}/members/{account_id}/$ref"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		{
			Name:    "GroupsService.ListMembers",
			Path:    []string{"/api/v0/groups/{id=*}/members/$ref"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for GroupsService service

type GroupsService interface {
	// Lists groups
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...client.CallOption) (*ListGroupsResponse, error)
	// Gets an groups
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...client.CallOption) (*v1.Group, error)
	// Creates a group
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*v1.Group, error)
	// Updates a group
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...client.CallOption) (*v1.Group, error)
	// Deletes a group
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...client.CallOption) (*emptypb.Empty, error)
	// group:addmember https://docs.microsoft.com/en-us/graph/api/group-post-members?view=graph-rest-1.0&tabs=http
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...client.CallOption) (*v1.Group, error)
	// group:removemember https://docs.microsoft.com/en-us/graph/api/group-delete-members?view=graph-rest-1.0
	RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...client.CallOption) (*v1.Group, error)
	// group:listmembers https://docs.microsoft.com/en-us/graph/api/group-list-members?view=graph-rest-1.0
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...client.CallOption) (*ListMembersResponse, error)
}

type groupsService struct {
	c    client.Client
	name string
}

func NewGroupsService(name string, c client.Client) GroupsService {
	return &groupsService{
		c:    c,
		name: name,
	}
}

func (c *groupsService) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...client.CallOption) (*ListGroupsResponse, error) {
	req := c.c.NewRequest(c.name, "GroupsService.ListGroups", in)
	out := new(ListGroupsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...client.CallOption) (*v1.Group, error) {
	req := c.c.NewRequest(c.name, "GroupsService.GetGroup", in)
	out := new(v1.Group)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*v1.Group, error) {
	req := c.c.NewRequest(c.name, "GroupsService.CreateGroup", in)
	out := new(v1.Group)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...client.CallOption) (*v1.Group, error) {
	req := c.c.NewRequest(c.name, "GroupsService.UpdateGroup", in)
	out := new(v1.Group)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "GroupsService.DeleteGroup", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) AddMember(ctx context.Context, in *AddMemberRequest, opts ...client.CallOption) (*v1.Group, error) {
	req := c.c.NewRequest(c.name, "GroupsService.AddMember", in)
	out := new(v1.Group)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...client.CallOption) (*v1.Group, error) {
	req := c.c.NewRequest(c.name, "GroupsService.RemoveMember", in)
	out := new(v1.Group)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsService) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...client.CallOption) (*ListMembersResponse, error) {
	req := c.c.NewRequest(c.name, "GroupsService.ListMembers", in)
	out := new(ListMembersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupsService service

type GroupsServiceHandler interface {
	// Lists groups
	ListGroups(context.Context, *ListGroupsRequest, *ListGroupsResponse) error
	// Gets an groups
	GetGroup(context.Context, *GetGroupRequest, *v1.Group) error
	// Creates a group
	CreateGroup(context.Context, *CreateGroupRequest, *v1.Group) error
	// Updates a group
	UpdateGroup(context.Context, *UpdateGroupRequest, *v1.Group) error
	// Deletes a group
	DeleteGroup(context.Context, *DeleteGroupRequest, *emptypb.Empty) error
	// group:addmember https://docs.microsoft.com/en-us/graph/api/group-post-members?view=graph-rest-1.0&tabs=http
	AddMember(context.Context, *AddMemberRequest, *v1.Group) error
	// group:removemember https://docs.microsoft.com/en-us/graph/api/group-delete-members?view=graph-rest-1.0
	RemoveMember(context.Context, *RemoveMemberRequest, *v1.Group) error
	// group:listmembers https://docs.microsoft.com/en-us/graph/api/group-list-members?view=graph-rest-1.0
	ListMembers(context.Context, *ListMembersRequest, *ListMembersResponse) error
}

func RegisterGroupsServiceHandler(s server.Server, hdlr GroupsServiceHandler, opts ...server.HandlerOption) error {
	type groupsService interface {
		ListGroups(ctx context.Context, in *ListGroupsRequest, out *ListGroupsResponse) error
		GetGroup(ctx context.Context, in *GetGroupRequest, out *v1.Group) error
		CreateGroup(ctx context.Context, in *CreateGroupRequest, out *v1.Group) error
		UpdateGroup(ctx context.Context, in *UpdateGroupRequest, out *v1.Group) error
		DeleteGroup(ctx context.Context, in *DeleteGroupRequest, out *emptypb.Empty) error
		AddMember(ctx context.Context, in *AddMemberRequest, out *v1.Group) error
		RemoveMember(ctx context.Context, in *RemoveMemberRequest, out *v1.Group) error
		ListMembers(ctx context.Context, in *ListMembersRequest, out *ListMembersResponse) error
	}
	type GroupsService struct {
		groupsService
	}
	h := &groupsServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.ListGroups",
		Path:    []string{"/api/v0/accounts/groups-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.GetGroup",
		Path:    []string{"/api/v0/accounts/groups-get"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.CreateGroup",
		Path:    []string{"/api/v0/accounts/groups-create"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.UpdateGroup",
		Path:    []string{"/api/v0/accounts/groups-update"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.DeleteGroup",
		Path:    []string{"/api/v0/accounts/groups-delete"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.AddMember",
		Path:    []string{"/api/v0/groups/{group_id=*}/members/$ref"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.RemoveMember",
		Path:    []string{"/api/v0/groups/{group_id=*}/members/{account_id}/$ref"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "GroupsService.ListMembers",
		Path:    []string{"/api/v0/groups/{id=*}/members/$ref"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&GroupsService{h}, opts...))
}

type groupsServiceHandler struct {
	GroupsServiceHandler
}

func (h *groupsServiceHandler) ListGroups(ctx context.Context, in *ListGroupsRequest, out *ListGroupsResponse) error {
	return h.GroupsServiceHandler.ListGroups(ctx, in, out)
}

func (h *groupsServiceHandler) GetGroup(ctx context.Context, in *GetGroupRequest, out *v1.Group) error {
	return h.GroupsServiceHandler.GetGroup(ctx, in, out)
}

func (h *groupsServiceHandler) CreateGroup(ctx context.Context, in *CreateGroupRequest, out *v1.Group) error {
	return h.GroupsServiceHandler.CreateGroup(ctx, in, out)
}

func (h *groupsServiceHandler) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, out *v1.Group) error {
	return h.GroupsServiceHandler.UpdateGroup(ctx, in, out)
}

func (h *groupsServiceHandler) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, out *emptypb.Empty) error {
	return h.GroupsServiceHandler.DeleteGroup(ctx, in, out)
}

func (h *groupsServiceHandler) AddMember(ctx context.Context, in *AddMemberRequest, out *v1.Group) error {
	return h.GroupsServiceHandler.AddMember(ctx, in, out)
}

func (h *groupsServiceHandler) RemoveMember(ctx context.Context, in *RemoveMemberRequest, out *v1.Group) error {
	return h.GroupsServiceHandler.RemoveMember(ctx, in, out)
}

func (h *groupsServiceHandler) ListMembers(ctx context.Context, in *ListMembersRequest, out *ListMembersResponse) error {
	return h.GroupsServiceHandler.ListMembers(ctx, in, out)
}

// Api Endpoints for IndexService service

func NewIndexServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "IndexService.RebuildIndex",
			Path:    []string{"/api/v0/index/rebuild"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for IndexService service

type IndexService interface {
	RebuildIndex(ctx context.Context, in *RebuildIndexRequest, opts ...client.CallOption) (*RebuildIndexResponse, error)
}

type indexService struct {
	c    client.Client
	name string
}

func NewIndexService(name string, c client.Client) IndexService {
	return &indexService{
		c:    c,
		name: name,
	}
}

func (c *indexService) RebuildIndex(ctx context.Context, in *RebuildIndexRequest, opts ...client.CallOption) (*RebuildIndexResponse, error) {
	req := c.c.NewRequest(c.name, "IndexService.RebuildIndex", in)
	out := new(RebuildIndexResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IndexService service

type IndexServiceHandler interface {
	RebuildIndex(context.Context, *RebuildIndexRequest, *RebuildIndexResponse) error
}

func RegisterIndexServiceHandler(s server.Server, hdlr IndexServiceHandler, opts ...server.HandlerOption) error {
	type indexService interface {
		RebuildIndex(ctx context.Context, in *RebuildIndexRequest, out *RebuildIndexResponse) error
	}
	type IndexService struct {
		indexService
	}
	h := &indexServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "IndexService.RebuildIndex",
		Path:    []string{"/api/v0/index/rebuild"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&IndexService{h}, opts...))
}

type indexServiceHandler struct {
	IndexServiceHandler
}

func (h *indexServiceHandler) RebuildIndex(ctx context.Context, in *RebuildIndexRequest, out *RebuildIndexResponse) error {
	return h.IndexServiceHandler.RebuildIndex(ctx, in, out)
}
