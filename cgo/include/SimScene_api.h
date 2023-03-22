#pragma once

#include "SimExport.h"
#include "SimForce_api.h"

typedef struct SimScene_T SimScene_T;

#ifdef __cplusplus
extern "C" {
#endif

	SIM_API SimScene_T* SimScene_Scene();
	SIM_API void SimScene_SetResourceDir(SimScene_T* scene, const char* dir);
	SIM_API char SimScene_LoadModel(SimScene_T* scene, const char* modelName);
	SIM_API unsigned int SimScene_AddForce(SimScene_T* scene, int entityId, unsigned int forceId);
	SIM_API unsigned int SimScene_AddRemoteForce(SimScene_T* scene, unsigned int forceId);
	SIM_API SimForce_T* SimScene_GetForce(SimScene_T* scene, unsigned int forceId);
	SIM_API void SimScene_RemoveForce(SimScene_T* scene, unsigned int forceId);

	SIM_API void SimScene_SetCalcFusion(SimScene_T* scene, char value);
	SIM_API void SimScene_MergePrepare(SimScene_T* scene);
	SIM_API void* SimScene_MergeData(SimScene_T* scene);
	SIM_API int SimScene_MergeDataLength(SimScene_T* scene);
	SIM_API void SimScene_SetMergeResult(SimScene_T* scene, const void* data, int size);

	SIM_API void SimScene_MessagePrepare(SimScene_T* scene);
	SIM_API void* SimScene_MessageData(SimScene_T* scene);
	SIM_API int SimScene_MessageDataLength(SimScene_T* scene);

	SIM_API void SimScene_ReceiveMessage(SimScene_T* scene, const void* data, int size);
	SIM_API void SimScene_ProcessMessageOnPostProcess(SimScene_T* scene);
	SIM_API void SimScene_ProcessMessageOnNextstep(SimScene_T* scene);

#ifdef __cplusplus
}
#endif