// Code generated by protoc-gen-go. DO NOT EDIT.
// source: responses.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Response that contains a result set.
type ResultSetResponse struct {
	ConnectionId string     `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	StatementId  uint32     `protobuf:"varint,2,opt,name=statement_id,json=statementId,proto3" json:"statement_id,omitempty"`
	OwnStatement bool       `protobuf:"varint,3,opt,name=own_statement,json=ownStatement,proto3" json:"own_statement,omitempty"`
	Signature    *Signature `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	FirstFrame   *Frame     `protobuf:"bytes,5,opt,name=first_frame,json=firstFrame,proto3" json:"first_frame,omitempty"`
	UpdateCount  uint64     `protobuf:"varint,6,opt,name=update_count,json=updateCount,proto3" json:"update_count,omitempty"`
	// with no signature nor other data.
	Metadata             *RpcMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResultSetResponse) Reset()         { *m = ResultSetResponse{} }
func (m *ResultSetResponse) String() string { return proto.CompactTextString(m) }
func (*ResultSetResponse) ProtoMessage()    {}
func (*ResultSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{0}
}
func (m *ResultSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultSetResponse.Unmarshal(m, b)
}
func (m *ResultSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultSetResponse.Marshal(b, m, deterministic)
}
func (dst *ResultSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultSetResponse.Merge(dst, src)
}
func (m *ResultSetResponse) XXX_Size() int {
	return xxx_messageInfo_ResultSetResponse.Size(m)
}
func (m *ResultSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResultSetResponse proto.InternalMessageInfo

func (m *ResultSetResponse) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *ResultSetResponse) GetStatementId() uint32 {
	if m != nil {
		return m.StatementId
	}
	return 0
}

func (m *ResultSetResponse) GetOwnStatement() bool {
	if m != nil {
		return m.OwnStatement
	}
	return false
}

func (m *ResultSetResponse) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *ResultSetResponse) GetFirstFrame() *Frame {
	if m != nil {
		return m.FirstFrame
	}
	return nil
}

func (m *ResultSetResponse) GetUpdateCount() uint64 {
	if m != nil {
		return m.UpdateCount
	}
	return 0
}

func (m *ResultSetResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to PrepareAndExecuteRequest
type ExecuteResponse struct {
	Results              []*ResultSetResponse `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	MissingStatement     bool                 `protobuf:"varint,2,opt,name=missing_statement,json=missingStatement,proto3" json:"missing_statement,omitempty"`
	Metadata             *RpcMetadata         `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExecuteResponse) Reset()         { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()    {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{1}
}
func (m *ExecuteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteResponse.Unmarshal(m, b)
}
func (m *ExecuteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteResponse.Marshal(b, m, deterministic)
}
func (dst *ExecuteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteResponse.Merge(dst, src)
}
func (m *ExecuteResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteResponse.Size(m)
}
func (m *ExecuteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteResponse proto.InternalMessageInfo

func (m *ExecuteResponse) GetResults() []*ResultSetResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ExecuteResponse) GetMissingStatement() bool {
	if m != nil {
		return m.MissingStatement
	}
	return false
}

func (m *ExecuteResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to PrepareRequest
type PrepareResponse struct {
	Statement            *StatementHandle `protobuf:"bytes,1,opt,name=statement,proto3" json:"statement,omitempty"`
	Metadata             *RpcMetadata     `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PrepareResponse) Reset()         { *m = PrepareResponse{} }
func (m *PrepareResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareResponse) ProtoMessage()    {}
func (*PrepareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{2}
}
func (m *PrepareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareResponse.Unmarshal(m, b)
}
func (m *PrepareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareResponse.Marshal(b, m, deterministic)
}
func (dst *PrepareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareResponse.Merge(dst, src)
}
func (m *PrepareResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareResponse.Size(m)
}
func (m *PrepareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareResponse proto.InternalMessageInfo

func (m *PrepareResponse) GetStatement() *StatementHandle {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *PrepareResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to FetchRequest
type FetchResponse struct {
	Frame                *Frame       `protobuf:"bytes,1,opt,name=frame,proto3" json:"frame,omitempty"`
	MissingStatement     bool         `protobuf:"varint,2,opt,name=missing_statement,json=missingStatement,proto3" json:"missing_statement,omitempty"`
	MissingResults       bool         `protobuf:"varint,3,opt,name=missing_results,json=missingResults,proto3" json:"missing_results,omitempty"`
	Metadata             *RpcMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FetchResponse) Reset()         { *m = FetchResponse{} }
func (m *FetchResponse) String() string { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()    {}
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{3}
}
func (m *FetchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchResponse.Unmarshal(m, b)
}
func (m *FetchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchResponse.Marshal(b, m, deterministic)
}
func (dst *FetchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchResponse.Merge(dst, src)
}
func (m *FetchResponse) XXX_Size() int {
	return xxx_messageInfo_FetchResponse.Size(m)
}
func (m *FetchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchResponse proto.InternalMessageInfo

func (m *FetchResponse) GetFrame() *Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *FetchResponse) GetMissingStatement() bool {
	if m != nil {
		return m.MissingStatement
	}
	return false
}

func (m *FetchResponse) GetMissingResults() bool {
	if m != nil {
		return m.MissingResults
	}
	return false
}

func (m *FetchResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to CreateStatementRequest
type CreateStatementResponse struct {
	ConnectionId         string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	StatementId          uint32       `protobuf:"varint,2,opt,name=statement_id,json=statementId,proto3" json:"statement_id,omitempty"`
	Metadata             *RpcMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateStatementResponse) Reset()         { *m = CreateStatementResponse{} }
func (m *CreateStatementResponse) String() string { return proto.CompactTextString(m) }
func (*CreateStatementResponse) ProtoMessage()    {}
func (*CreateStatementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{4}
}
func (m *CreateStatementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStatementResponse.Unmarshal(m, b)
}
func (m *CreateStatementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStatementResponse.Marshal(b, m, deterministic)
}
func (dst *CreateStatementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStatementResponse.Merge(dst, src)
}
func (m *CreateStatementResponse) XXX_Size() int {
	return xxx_messageInfo_CreateStatementResponse.Size(m)
}
func (m *CreateStatementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStatementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStatementResponse proto.InternalMessageInfo

func (m *CreateStatementResponse) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *CreateStatementResponse) GetStatementId() uint32 {
	if m != nil {
		return m.StatementId
	}
	return 0
}

func (m *CreateStatementResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to CloseStatementRequest
type CloseStatementResponse struct {
	Metadata             *RpcMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CloseStatementResponse) Reset()         { *m = CloseStatementResponse{} }
func (m *CloseStatementResponse) String() string { return proto.CompactTextString(m) }
func (*CloseStatementResponse) ProtoMessage()    {}
func (*CloseStatementResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{5}
}
func (m *CloseStatementResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseStatementResponse.Unmarshal(m, b)
}
func (m *CloseStatementResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseStatementResponse.Marshal(b, m, deterministic)
}
func (dst *CloseStatementResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseStatementResponse.Merge(dst, src)
}
func (m *CloseStatementResponse) XXX_Size() int {
	return xxx_messageInfo_CloseStatementResponse.Size(m)
}
func (m *CloseStatementResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseStatementResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseStatementResponse proto.InternalMessageInfo

func (m *CloseStatementResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to OpenConnectionRequest {
type OpenConnectionResponse struct {
	Metadata             *RpcMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OpenConnectionResponse) Reset()         { *m = OpenConnectionResponse{} }
func (m *OpenConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*OpenConnectionResponse) ProtoMessage()    {}
func (*OpenConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{6}
}
func (m *OpenConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenConnectionResponse.Unmarshal(m, b)
}
func (m *OpenConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenConnectionResponse.Marshal(b, m, deterministic)
}
func (dst *OpenConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenConnectionResponse.Merge(dst, src)
}
func (m *OpenConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_OpenConnectionResponse.Size(m)
}
func (m *OpenConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenConnectionResponse proto.InternalMessageInfo

func (m *OpenConnectionResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to CloseConnectionRequest {
type CloseConnectionResponse struct {
	Metadata             *RpcMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CloseConnectionResponse) Reset()         { *m = CloseConnectionResponse{} }
func (m *CloseConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*CloseConnectionResponse) ProtoMessage()    {}
func (*CloseConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{7}
}
func (m *CloseConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseConnectionResponse.Unmarshal(m, b)
}
func (m *CloseConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseConnectionResponse.Marshal(b, m, deterministic)
}
func (dst *CloseConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseConnectionResponse.Merge(dst, src)
}
func (m *CloseConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_CloseConnectionResponse.Size(m)
}
func (m *CloseConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseConnectionResponse proto.InternalMessageInfo

func (m *CloseConnectionResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to ConnectionSyncRequest
type ConnectionSyncResponse struct {
	ConnProps            *ConnectionProperties `protobuf:"bytes,1,opt,name=conn_props,json=connProps,proto3" json:"conn_props,omitempty"`
	Metadata             *RpcMetadata          `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConnectionSyncResponse) Reset()         { *m = ConnectionSyncResponse{} }
func (m *ConnectionSyncResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectionSyncResponse) ProtoMessage()    {}
func (*ConnectionSyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{8}
}
func (m *ConnectionSyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionSyncResponse.Unmarshal(m, b)
}
func (m *ConnectionSyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionSyncResponse.Marshal(b, m, deterministic)
}
func (dst *ConnectionSyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionSyncResponse.Merge(dst, src)
}
func (m *ConnectionSyncResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectionSyncResponse.Size(m)
}
func (m *ConnectionSyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionSyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionSyncResponse proto.InternalMessageInfo

func (m *ConnectionSyncResponse) GetConnProps() *ConnectionProperties {
	if m != nil {
		return m.ConnProps
	}
	return nil
}

func (m *ConnectionSyncResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type DatabasePropertyElement struct {
	Key                  *DatabaseProperty `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *TypedValue       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Metadata             *RpcMetadata      `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DatabasePropertyElement) Reset()         { *m = DatabasePropertyElement{} }
func (m *DatabasePropertyElement) String() string { return proto.CompactTextString(m) }
func (*DatabasePropertyElement) ProtoMessage()    {}
func (*DatabasePropertyElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{9}
}
func (m *DatabasePropertyElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabasePropertyElement.Unmarshal(m, b)
}
func (m *DatabasePropertyElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabasePropertyElement.Marshal(b, m, deterministic)
}
func (dst *DatabasePropertyElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabasePropertyElement.Merge(dst, src)
}
func (m *DatabasePropertyElement) XXX_Size() int {
	return xxx_messageInfo_DatabasePropertyElement.Size(m)
}
func (m *DatabasePropertyElement) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabasePropertyElement.DiscardUnknown(m)
}

var xxx_messageInfo_DatabasePropertyElement proto.InternalMessageInfo

func (m *DatabasePropertyElement) GetKey() *DatabaseProperty {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *DatabasePropertyElement) GetValue() *TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DatabasePropertyElement) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response for Meta#getDatabaseProperties()
type DatabasePropertyResponse struct {
	Props                []*DatabasePropertyElement `protobuf:"bytes,1,rep,name=props,proto3" json:"props,omitempty"`
	Metadata             *RpcMetadata               `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *DatabasePropertyResponse) Reset()         { *m = DatabasePropertyResponse{} }
func (m *DatabasePropertyResponse) String() string { return proto.CompactTextString(m) }
func (*DatabasePropertyResponse) ProtoMessage()    {}
func (*DatabasePropertyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{10}
}
func (m *DatabasePropertyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabasePropertyResponse.Unmarshal(m, b)
}
func (m *DatabasePropertyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabasePropertyResponse.Marshal(b, m, deterministic)
}
func (dst *DatabasePropertyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabasePropertyResponse.Merge(dst, src)
}
func (m *DatabasePropertyResponse) XXX_Size() int {
	return xxx_messageInfo_DatabasePropertyResponse.Size(m)
}
func (m *DatabasePropertyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabasePropertyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatabasePropertyResponse proto.InternalMessageInfo

func (m *DatabasePropertyResponse) GetProps() []*DatabasePropertyElement {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *DatabasePropertyResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Send contextual information about some error over the wire from the server.
type ErrorResponse struct {
	Exceptions           []string     `protobuf:"bytes,1,rep,name=exceptions,proto3" json:"exceptions,omitempty"`
	HasExceptions        bool         `protobuf:"varint,7,opt,name=has_exceptions,json=hasExceptions,proto3" json:"has_exceptions,omitempty"`
	ErrorMessage         string       `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Severity             Severity     `protobuf:"varint,3,opt,name=severity,proto3,enum=Severity" json:"severity,omitempty"`
	ErrorCode            uint32       `protobuf:"varint,4,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	SqlState             string       `protobuf:"bytes,5,opt,name=sql_state,json=sqlState,proto3" json:"sql_state,omitempty"`
	Metadata             *RpcMetadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{11}
}
func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (dst *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(dst, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetExceptions() []string {
	if m != nil {
		return m.Exceptions
	}
	return nil
}

func (m *ErrorResponse) GetHasExceptions() bool {
	if m != nil {
		return m.HasExceptions
	}
	return false
}

func (m *ErrorResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *ErrorResponse) GetSeverity() Severity {
	if m != nil {
		return m.Severity
	}
	return Severity_UNKNOWN_SEVERITY
}

func (m *ErrorResponse) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *ErrorResponse) GetSqlState() string {
	if m != nil {
		return m.SqlState
	}
	return ""
}

func (m *ErrorResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type SyncResultsResponse struct {
	MissingStatement     bool         `protobuf:"varint,1,opt,name=missing_statement,json=missingStatement,proto3" json:"missing_statement,omitempty"`
	MoreResults          bool         `protobuf:"varint,2,opt,name=more_results,json=moreResults,proto3" json:"more_results,omitempty"`
	Metadata             *RpcMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SyncResultsResponse) Reset()         { *m = SyncResultsResponse{} }
func (m *SyncResultsResponse) String() string { return proto.CompactTextString(m) }
func (*SyncResultsResponse) ProtoMessage()    {}
func (*SyncResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{12}
}
func (m *SyncResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncResultsResponse.Unmarshal(m, b)
}
func (m *SyncResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncResultsResponse.Marshal(b, m, deterministic)
}
func (dst *SyncResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncResultsResponse.Merge(dst, src)
}
func (m *SyncResultsResponse) XXX_Size() int {
	return xxx_messageInfo_SyncResultsResponse.Size(m)
}
func (m *SyncResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncResultsResponse proto.InternalMessageInfo

func (m *SyncResultsResponse) GetMissingStatement() bool {
	if m != nil {
		return m.MissingStatement
	}
	return false
}

func (m *SyncResultsResponse) GetMoreResults() bool {
	if m != nil {
		return m.MoreResults
	}
	return false
}

func (m *SyncResultsResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Generic metadata for the server to return with each response.
type RpcMetadata struct {
	ServerAddress        string   `protobuf:"bytes,1,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcMetadata) Reset()         { *m = RpcMetadata{} }
func (m *RpcMetadata) String() string { return proto.CompactTextString(m) }
func (*RpcMetadata) ProtoMessage()    {}
func (*RpcMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{13}
}
func (m *RpcMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcMetadata.Unmarshal(m, b)
}
func (m *RpcMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcMetadata.Marshal(b, m, deterministic)
}
func (dst *RpcMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcMetadata.Merge(dst, src)
}
func (m *RpcMetadata) XXX_Size() int {
	return xxx_messageInfo_RpcMetadata.Size(m)
}
func (m *RpcMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RpcMetadata proto.InternalMessageInfo

func (m *RpcMetadata) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

// Response to a commit request
type CommitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitResponse) Reset()         { *m = CommitResponse{} }
func (m *CommitResponse) String() string { return proto.CompactTextString(m) }
func (*CommitResponse) ProtoMessage()    {}
func (*CommitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{14}
}
func (m *CommitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitResponse.Unmarshal(m, b)
}
func (m *CommitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitResponse.Marshal(b, m, deterministic)
}
func (dst *CommitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitResponse.Merge(dst, src)
}
func (m *CommitResponse) XXX_Size() int {
	return xxx_messageInfo_CommitResponse.Size(m)
}
func (m *CommitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommitResponse proto.InternalMessageInfo

// Response to a rollback request
type RollbackResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RollbackResponse) Reset()         { *m = RollbackResponse{} }
func (m *RollbackResponse) String() string { return proto.CompactTextString(m) }
func (*RollbackResponse) ProtoMessage()    {}
func (*RollbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{15}
}
func (m *RollbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollbackResponse.Unmarshal(m, b)
}
func (m *RollbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollbackResponse.Marshal(b, m, deterministic)
}
func (dst *RollbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollbackResponse.Merge(dst, src)
}
func (m *RollbackResponse) XXX_Size() int {
	return xxx_messageInfo_RollbackResponse.Size(m)
}
func (m *RollbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RollbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RollbackResponse proto.InternalMessageInfo

// Response to a batch update request
type ExecuteBatchResponse struct {
	ConnectionId         string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	StatementId          uint32       `protobuf:"varint,2,opt,name=statement_id,json=statementId,proto3" json:"statement_id,omitempty"`
	UpdateCounts         []uint64     `protobuf:"varint,3,rep,packed,name=update_counts,json=updateCounts,proto3" json:"update_counts,omitempty"`
	MissingStatement     bool         `protobuf:"varint,4,opt,name=missing_statement,json=missingStatement,proto3" json:"missing_statement,omitempty"`
	Metadata             *RpcMetadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ExecuteBatchResponse) Reset()         { *m = ExecuteBatchResponse{} }
func (m *ExecuteBatchResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteBatchResponse) ProtoMessage()    {}
func (*ExecuteBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_responses_d38978c60cf7d175, []int{16}
}
func (m *ExecuteBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteBatchResponse.Unmarshal(m, b)
}
func (m *ExecuteBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteBatchResponse.Marshal(b, m, deterministic)
}
func (dst *ExecuteBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteBatchResponse.Merge(dst, src)
}
func (m *ExecuteBatchResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteBatchResponse.Size(m)
}
func (m *ExecuteBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteBatchResponse proto.InternalMessageInfo

func (m *ExecuteBatchResponse) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *ExecuteBatchResponse) GetStatementId() uint32 {
	if m != nil {
		return m.StatementId
	}
	return 0
}

func (m *ExecuteBatchResponse) GetUpdateCounts() []uint64 {
	if m != nil {
		return m.UpdateCounts
	}
	return nil
}

func (m *ExecuteBatchResponse) GetMissingStatement() bool {
	if m != nil {
		return m.MissingStatement
	}
	return false
}

func (m *ExecuteBatchResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ResultSetResponse)(nil), "ResultSetResponse")
	proto.RegisterType((*ExecuteResponse)(nil), "ExecuteResponse")
	proto.RegisterType((*PrepareResponse)(nil), "PrepareResponse")
	proto.RegisterType((*FetchResponse)(nil), "FetchResponse")
	proto.RegisterType((*CreateStatementResponse)(nil), "CreateStatementResponse")
	proto.RegisterType((*CloseStatementResponse)(nil), "CloseStatementResponse")
	proto.RegisterType((*OpenConnectionResponse)(nil), "OpenConnectionResponse")
	proto.RegisterType((*CloseConnectionResponse)(nil), "CloseConnectionResponse")
	proto.RegisterType((*ConnectionSyncResponse)(nil), "ConnectionSyncResponse")
	proto.RegisterType((*DatabasePropertyElement)(nil), "DatabasePropertyElement")
	proto.RegisterType((*DatabasePropertyResponse)(nil), "DatabasePropertyResponse")
	proto.RegisterType((*ErrorResponse)(nil), "ErrorResponse")
	proto.RegisterType((*SyncResultsResponse)(nil), "SyncResultsResponse")
	proto.RegisterType((*RpcMetadata)(nil), "RpcMetadata")
	proto.RegisterType((*CommitResponse)(nil), "CommitResponse")
	proto.RegisterType((*RollbackResponse)(nil), "RollbackResponse")
	proto.RegisterType((*ExecuteBatchResponse)(nil), "ExecuteBatchResponse")
}

func init() { proto.RegisterFile("responses.proto", fileDescriptor_responses_d38978c60cf7d175) }

var fileDescriptor_responses_d38978c60cf7d175 = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xc6, 0xea, 0xc7, 0x36, 0x87, 0xa4, 0x2d, 0xb3, 0x6d, 0x42, 0xf4, 0x0f, 0xb2, 0x8c, 0x20,
	0x04, 0x5a, 0xf0, 0xe0, 0xe6, 0x05, 0x6a, 0x55, 0x41, 0x73, 0x08, 0x1a, 0xac, 0x8a, 0x5e, 0x89,
	0x35, 0x39, 0x91, 0x08, 0x93, 0x5c, 0x66, 0x77, 0xa5, 0x58, 0x6f, 0x50, 0xa0, 0x87, 0xde, 0x7a,
	0xee, 0x0b, 0xf4, 0x85, 0xfa, 0x34, 0xc5, 0x72, 0x29, 0x92, 0x69, 0x64, 0x35, 0x6a, 0x7c, 0x92,
	0xf8, 0xcd, 0xf0, 0x9b, 0xd1, 0xb7, 0xb3, 0xdf, 0x08, 0xce, 0x04, 0xca, 0x92, 0x17, 0x12, 0x65,
	0x58, 0x0a, 0xae, 0xf8, 0xe7, 0x4e, 0xcc, 0xf3, 0x9c, 0x17, 0xe6, 0x69, 0xf2, 0x67, 0x0f, 0xce,
	0x29, 0xca, 0x55, 0xa6, 0xe6, 0xa8, 0x68, 0x9d, 0xea, 0x5d, 0x82, 0x1b, 0xf3, 0xa2, 0xc0, 0x58,
	0xa5, 0xbc, 0x88, 0xd2, 0xc4, 0x27, 0x63, 0x12, 0x58, 0xd4, 0x69, 0xc1, 0x17, 0x89, 0x77, 0x01,
	0x8e, 0x54, 0x4c, 0x61, 0x8e, 0x85, 0xd2, 0x39, 0xbd, 0x31, 0x09, 0x5c, 0x6a, 0x37, 0xd8, 0x8b,
	0x44, 0xf3, 0xf0, 0xb7, 0x45, 0xd4, 0x40, 0x7e, 0x7f, 0x4c, 0x82, 0x13, 0xea, 0xf0, 0xb7, 0xc5,
	0x7c, 0x8b, 0x79, 0x01, 0x58, 0x32, 0x5d, 0x14, 0x4c, 0xad, 0x04, 0xfa, 0x83, 0x31, 0x09, 0xec,
	0x2b, 0x08, 0xe7, 0x5b, 0x84, 0xb6, 0x41, 0xef, 0x29, 0xd8, 0xaf, 0x53, 0x21, 0x55, 0xf4, 0x5a,
	0xb0, 0x1c, 0xfd, 0x61, 0x95, 0x7b, 0x14, 0x3e, 0xd7, 0x4f, 0x14, 0xaa, 0x50, 0xf5, 0x5d, 0xb7,
	0xb6, 0x2a, 0x13, 0xa6, 0x30, 0x8a, 0xf9, 0xaa, 0x50, 0xfe, 0xd1, 0x98, 0x04, 0x03, 0x6a, 0x1b,
	0x6c, 0xaa, 0x21, 0x2f, 0x80, 0x93, 0x1c, 0x15, 0x4b, 0x98, 0x62, 0xfe, 0x71, 0x45, 0xe4, 0x84,
	0xb4, 0x8c, 0x5f, 0xd6, 0x18, 0x6d, 0xa2, 0x93, 0x3f, 0x08, 0x9c, 0xcd, 0xee, 0x30, 0x5e, 0x29,
	0x6c, 0x04, 0xfa, 0x16, 0x8e, 0x45, 0xa5, 0x9a, 0xf4, 0xc9, 0xb8, 0x1f, 0xd8, 0x57, 0x5e, 0xf8,
	0x9e, 0x8a, 0x74, 0x9b, 0xe2, 0x7d, 0x03, 0xe7, 0x79, 0x2a, 0x65, 0x5a, 0x2c, 0x3a, 0x52, 0xf4,
	0x2a, 0x29, 0x46, 0x75, 0xa0, 0x2b, 0x47, 0xdb, 0x58, 0x7f, 0x6f, 0x63, 0xb7, 0x70, 0xf6, 0x4a,
	0x60, 0xc9, 0x44, 0xdb, 0x57, 0x08, 0x56, 0x5b, 0x81, 0x54, 0x6f, 0x8f, 0xc2, 0x86, 0xfb, 0x47,
	0x56, 0x24, 0x99, 0x56, 0x74, 0x67, 0xb1, 0xde, 0xde, 0x62, 0x7f, 0x11, 0x70, 0x9f, 0xa3, 0x8a,
	0x97, 0x4d, 0xad, 0x2f, 0x61, 0x68, 0xce, 0x81, 0xbc, 0x73, 0x0e, 0x06, 0x3c, 0xec, 0x37, 0x3f,
	0x85, 0xb3, 0x6d, 0xf2, 0x56, 0x56, 0x33, 0x29, 0xa7, 0x35, 0x4c, 0x6b, 0x25, 0xbb, 0xfd, 0x0e,
	0xf6, 0xf6, 0xfb, 0x1b, 0x81, 0xc7, 0x53, 0x81, 0x4c, 0x61, 0x53, 0xe6, 0xc1, 0xc7, 0xfb, 0xc3,
	0x8f, 0xea, 0x1a, 0x1e, 0x4d, 0x33, 0x2e, 0x77, 0xf4, 0xd2, 0xe5, 0x20, 0xff, 0xc5, 0xf1, 0x53,
	0x89, 0xc5, 0xb4, 0x69, 0xf2, 0x7f, 0x70, 0x4c, 0xe1, 0x71, 0xd5, 0xc7, 0x47, 0x91, 0xdc, 0xc1,
	0xa3, 0xf6, 0xfd, 0xf9, 0xa6, 0x88, 0x1b, 0x8e, 0x67, 0x00, 0x5a, 0xc3, 0xa8, 0x14, 0xbc, 0x94,
	0x35, 0xcb, 0x67, 0x61, 0x9b, 0xfc, 0x4a, 0xf0, 0x12, 0x85, 0x4a, 0x51, 0x52, 0x4b, 0x27, 0xea,
	0x67, 0x79, 0xc0, 0x10, 0xea, 0x43, 0xfd, 0x81, 0x29, 0x76, 0xc3, 0x24, 0xd6, 0x5c, 0x9b, 0x59,
	0x66, 0x66, 0xe8, 0x12, 0xfa, 0xb7, 0xb8, 0xa9, 0x8b, 0x9e, 0x87, 0xff, 0x4e, 0xa3, 0x3a, 0xea,
	0x5d, 0xc0, 0x70, 0xcd, 0xb2, 0x15, 0xd6, 0x75, 0xec, 0xf0, 0xe7, 0x4d, 0x89, 0xc9, 0x2f, 0x1a,
	0xa2, 0x26, 0x72, 0xc0, 0xa1, 0x2a, 0xf0, 0xdf, 0xab, 0xd2, 0x5e, 0xc4, 0xe1, 0x56, 0x04, 0x6d,
	0x0f, 0x7e, 0x78, 0x4f, 0xdb, 0xd4, 0xa4, 0x1d, 0xa0, 0xc1, 0xaf, 0x3d, 0x70, 0x67, 0x42, 0x70,
	0xd1, 0xd4, 0xfa, 0x1a, 0x00, 0xef, 0x62, 0x2c, 0xb5, 0xc2, 0xa6, 0xa0, 0x45, 0x3b, 0x88, 0xf7,
	0x04, 0x4e, 0x97, 0x4c, 0x46, 0x9d, 0x9c, 0xe3, 0xea, 0x72, 0xb9, 0x4b, 0x26, 0x67, 0x6d, 0xda,
	0x25, 0xb8, 0xa8, 0x79, 0xa3, 0x1c, 0xa5, 0x64, 0x0b, 0xa3, 0x91, 0x45, 0x9d, 0x0a, 0x7c, 0x69,
	0x30, 0xef, 0x09, 0x9c, 0x48, 0x5c, 0xa3, 0x48, 0xd5, 0xa6, 0x52, 0xe7, 0xf4, 0xca, 0x0a, 0xe7,
	0x35, 0x40, 0x9b, 0x90, 0xf7, 0x15, 0x80, 0xe1, 0x8a, 0x79, 0x62, 0x4c, 0xdd, 0xa5, 0x56, 0x85,
	0x4c, 0x79, 0x82, 0xde, 0x17, 0x60, 0xc9, 0x37, 0x99, 0x31, 0x86, 0xca, 0xc6, 0x2d, 0x7a, 0x22,
	0xdf, 0x64, 0xd5, 0xed, 0x78, 0x47, 0x8a, 0xa3, 0xbd, 0x52, 0xfc, 0x4e, 0xe0, 0x93, 0x7a, 0xfe,
	0xb4, 0x3b, 0x34, 0x82, 0xec, 0xf4, 0x1e, 0x72, 0x8f, 0xf7, 0x5c, 0x80, 0x93, 0x73, 0x81, 0x8d,
	0xf1, 0x18, 0x8f, 0xb2, 0x35, 0xb6, 0xcb, 0x75, 0xf6, 0x8f, 0xc4, 0x33, 0xb0, 0x3b, 0x01, 0xad,
	0xbc, 0x44, 0xb1, 0x46, 0x11, 0xb1, 0x24, 0x11, 0x28, 0x65, 0xed, 0x34, 0xae, 0x41, 0xbf, 0x37,
	0xe0, 0x64, 0x04, 0xa7, 0x53, 0x9e, 0xe7, 0x69, 0xe3, 0x0a, 0x13, 0x0f, 0x46, 0x94, 0x67, 0xd9,
	0x0d, 0x8b, 0x6f, 0x1b, 0xec, 0x6f, 0x02, 0x9f, 0xd6, 0x7b, 0xe8, 0x9a, 0x75, 0x8d, 0xf8, 0x01,
	0xb7, 0x75, 0x77, 0x6b, 0x6a, 0x0f, 0xee, 0x07, 0x03, 0xea, 0x74, 0xd6, 0xe6, 0x3d, 0xbb, 0x6c,
	0xf0, 0x01, 0xbb, 0x6c, 0xb8, 0x4f, 0xb8, 0xeb, 0x09, 0x8c, 0xb9, 0x58, 0x84, 0xac, 0x64, 0xf1,
	0x12, 0xc3, 0x98, 0x65, 0x71, 0xaa, 0x30, 0x64, 0x6b, 0xa6, 0xd2, 0x98, 0x99, 0xff, 0x2a, 0x37,
	0x47, 0xd5, 0xc7, 0x77, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0x76, 0x9a, 0x2c, 0xd3, 0x08,
	0x00, 0x00,
}
