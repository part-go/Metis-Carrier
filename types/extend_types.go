package types

import (
	pb "github.com/RosettaFlow/Carrier-Go/lib/api"
	apicommonpb "github.com/RosettaFlow/Carrier-Go/lib/common"
	libtypes "github.com/RosettaFlow/Carrier-Go/lib/types"
)

func NewTaskDetailShowFromTaskData(input *Task, role apicommonpb.TaskRole) *TaskEventShowAndRole {
	taskData := input.GetTaskData()
	detailShow := &TaskEventShowAndRole{
		Data: &pb.TaskDetailShow{
			TaskId:   taskData.GetTaskId(),
			TaskName: taskData.GetTaskName(),
			UserType: taskData.GetUserType(),
			User:     taskData.GetUser(),
			//TODO: 需要确认部分
			//Role:     role,
			Sender: &apicommonpb.TaskOrganization{
				PartyId:    taskData.GetPartyId(),
				NodeName:   taskData.GetNodeName(),
				NodeId:     taskData.GetNodeId(),
				IdentityId: taskData.GetIdentityId(),
			},
			AlgoSupplier: &apicommonpb.TaskOrganization{
				PartyId:    taskData.GetPartyId(),
				NodeName:   taskData.GetNodeName(),
				NodeId:     taskData.GetNodeId(),
				IdentityId: taskData.GetIdentityId(),
			},
			DataSuppliers:  make([]*pb.TaskDataSupplierShow, 0, len(taskData.GetDataSuppliers())),
			PowerSuppliers: make([]*pb.TaskPowerSupplierShow, 0, len(taskData.GetPowerSuppliers())),
			Receivers:      taskData.GetReceivers(),
			CreateAt:       taskData.GetCreateAt(),
			StartAt:        taskData.GetStartAt(),
			EndAt:          taskData.GetEndAt(),
			State:          taskData.GetState(),
			OperationCost: &apicommonpb.TaskResourceCostDeclare{
				Processor: taskData.GetOperationCost().GetProcessor(),
				Memory:    taskData.GetOperationCost().GetMemory(),
				Bandwidth: taskData.GetOperationCost().GetBandwidth(),
				Duration:  taskData.GetOperationCost().GetDuration(),
			},
		},
	}
	// DataSupplier
	for _, metadataSupplier := range taskData.GetDataSuppliers() {
		dataSupplier := &pb.TaskDataSupplierShow{
			Organization: &apicommonpb.TaskOrganization{
				PartyId:    metadataSupplier.GetOrganization().GetPartyId(),
				NodeName:   metadataSupplier.GetOrganization().GetNodeName(),
				NodeId:     metadataSupplier.GetOrganization().GetNodeId(),
				IdentityId: metadataSupplier.GetOrganization().GetIdentityId(),
			},
			MetadataId:   metadataSupplier.GetMetadataId(),
			MetadataName: metadataSupplier.GetMetadataName(),
		}
		detailShow.Data.DataSuppliers = append(detailShow.Data.DataSuppliers, dataSupplier)
	}
	// powerSupplier
	for _, data := range taskData.GetPowerSuppliers() {
		detailShow.Data.PowerSuppliers = append(detailShow.Data.PowerSuppliers, &pb.TaskPowerSupplierShow{
			Organization: &apicommonpb.TaskOrganization{
				PartyId:    data.GetOrganization().GetPartyId(),
				NodeName:   data.GetOrganization().GetNodeName(),
				NodeId:     data.GetOrganization().GetNodeId(),
				IdentityId: data.GetOrganization().GetIdentityId(),
			},
			PowerInfo: &libtypes.ResourceUsageOverview{
				TotalMem:       data.GetResourceUsedOverview().GetTotalMem(),
				UsedMem:        data.GetResourceUsedOverview().GetUsedMem(),
				TotalProcessor: data.GetResourceUsedOverview().GetTotalProcessor(),
				UsedProcessor:  data.GetResourceUsedOverview().GetUsedProcessor(),
				TotalBandwidth: data.GetResourceUsedOverview().GetTotalBandwidth(),
				UsedBandwidth:  data.GetResourceUsedOverview().GetUsedBandwidth(),
			},
		})
	}
	detailShow.Role = role
	return detailShow
}

func NewTaskEventFromAPIEvent(input []*libtypes.TaskEvent) []*pb.TaskEventShow {
	result := make([]*pb.TaskEventShow, 0, len(input))
	for _, event := range input {
		result = append(result, &pb.TaskEventShow{
			TaskId:   event.GetTaskId(),
			Type:     event.GetType(),
			CreateAt: event.GetCreateAt(),
			Content:  event.GetContent(),
			Owner: &apicommonpb.Organization{
				IdentityId: event.GetIdentityId(),
			},
		})
	}
	return result
}

func NewTotalMetadataInfoFromMetadata(input *Metadata) *pb.GetTotalMetadataDetailResponse {
	response := &pb.GetTotalMetadataDetailResponse{
		Owner: &apicommonpb.Organization{
			NodeName:   input.data.GetNodeName(),
			NodeId:     input.data.GetNodeId(),
			IdentityId: input.data.GetIdentityId(),
		},
		Information: &libtypes.MetadataDetail{
			MetadataSummary: &libtypes.MetadataSummary{
				MetadataId: input.GetData().GetDataId(),
				OriginId:   input.GetData().GetOriginId(),
				TableName:  input.GetData().GetTableName(),
				Desc:       input.GetData().GetDesc(),
				FilePath:   input.GetData().GetFilePath(),
				Rows:       input.GetData().GetRows(),
				Columns:    input.GetData().GetColumns(),
				Size_:      input.GetData().GetSize_(),
				FileType:   input.GetData().GetFileType(),
				HasTitle:   input.GetData().GetHasTitle(),
				Industry:   input.GetData().GetIndustry(),
				State:      input.GetData().GetState(),
			},
			MetadataColumns: input.GetData().GetMetadataColumns(),
		},
	}
	return response
}

func NewSelfMetadataInfoFromMetadata(islocal bool, input *Metadata) *pb.GetSelfMetadataDetailResponse {
	response := &pb.GetSelfMetadataDetailResponse{
		Owner: &apicommonpb.Organization{
			NodeName:   input.data.GetNodeName(),
			NodeId:     input.data.GetNodeId(),
			IdentityId: input.data.GetIdentityId(),
		},
		Information: &libtypes.MetadataDetail{
			MetadataSummary: &libtypes.MetadataSummary{
				MetadataId: input.GetData().GetDataId(),
				OriginId:   input.GetData().GetOriginId(),
				TableName:  input.GetData().GetTableName(),
				Desc:       input.GetData().GetDesc(),
				FilePath:   input.GetData().GetFilePath(),
				Rows:       input.GetData().GetRows(),
				Columns:    input.GetData().GetColumns(),
				Size_:      input.GetData().GetSize_(),
				FileType:   input.GetData().GetFileType(),
				HasTitle:   input.GetData().GetHasTitle(),
				Industry:   input.GetData().GetIndustry(),
				State:      input.GetData().GetState(),
			},
			MetadataColumns: input.GetData().GetMetadataColumns(),
		},
		IsLocal: islocal,
	}
	return response
}

func NewTotalMetadataInfoArrayFromMetadataArray(input MetadataArray) []*pb.GetTotalMetadataDetailResponse {
	result := make([]*pb.GetTotalMetadataDetailResponse, 0, input.Len())
	for _, metadata := range input {
		if metadata == nil {
			continue
		}
		result = append(result, NewTotalMetadataInfoFromMetadata(metadata))
	}
	return result
}

func NewSelfMetadataInfoArrayFromMetadataArray(localArr, publishArr MetadataArray) []*pb.GetSelfMetadataDetailResponse {
	result := make([]*pb.GetSelfMetadataDetailResponse, 0, localArr.Len()+publishArr.Len())

	for _, metadata := range localArr {
		if metadata == nil {
			continue
		}
		result = append(result, NewSelfMetadataInfoFromMetadata(true, metadata))
	}

	for _, metadata := range publishArr {
		if metadata == nil {
			continue
		}
		result = append(result, NewSelfMetadataInfoFromMetadata(false, metadata))
	}

	return result
}

func NewOrgResourceFromResource(input *Resource) *RemoteResourceTable {
	return &RemoteResourceTable{
		identityId: input.data.GetIdentityId(),
		total: &resource{
			mem:       input.data.GetTotalMem(),
			processor: input.data.GetTotalProcessor(),
			bandwidth: input.data.GetTotalBandwidth(),
		},
		used: &resource{
			mem:       input.data.GetUsedMem(),
			processor: input.data.GetUsedProcessor(),
			bandwidth: input.data.GetUsedBandwidth(),
		},
	}
}

//func NewOrgResourceArrayFromResourceArray(input ResourceArray) []*RemoteResourceTable {
//	result := make([]*RemoteResourceTable, input.Len())
//	for i, resource := range input {
//		result[i] = NewOrgResourceFromResource(resource)
//	}
//	return result
//}
