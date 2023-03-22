#pragma once

#include "SimExport.h"
#include "SimData_cgo.h"

typedef struct SimFusion_T SimFusion_T;

#ifdef __cplusplus
extern "C" {
#endif

	SIM_API SimFusion_T* SimFusion_New();
	SIM_API void SimFusion_ToMerge(SimFusion_T* fusion, const void* data, int len);
	SIM_API void SimFusion_OutMerge(SimFusion_T* fusion);
	SIM_API void* SimFusion_GetOutMergeData(SimFusion_T* fusion);
	SIM_API int SimFusion_GetOutMergeLength(SimFusion_T* fusion);
	SIM_API void SimFusion_NextStepPrepare(SimFusion_T* fusion);
	SIM_API void SimFusion_NextStep_Red(SimFusion_T* fusion, long long simTime, long long step);
	SIM_API void SimFusion_NextStep_Blue(SimFusion_T* fusion, long long simTime, long long step);
	SIM_API void SimFusion_NextStepFinished(SimFusion_T* fusion);

	SIM_API void SimFusion_AddForce(SimFusion_T* fusion, unsigned int id);
	SIM_API void SimFusion_RemoveForce(SimFusion_T* fusion, unsigned int id);

	SIM_API void SimFusion_SetForcePlatformCode(SimFusion_T* fusion, unsigned int id, unsigned int code);
	SIM_API void SimFusion_SetForcePlatformType(SimFusion_T* fusion, unsigned int id, const char* type);
	SIM_API void SimFusion_SetForceName(SimFusion_T* fusion, unsigned int id, const char* name);
	SIM_API void SimFusion_SetForceTeam(SimFusion_T* fusion, unsigned int id, const char* team);
	SIM_API void SimFusion_SetForcePosture(SimFusion_T* fusion, const ForcePosture params);

	SIM_API void SimFusion_PrepareDetected(SimFusion_T* fusion);
	SIM_API DetectedPostore* SimFusion_GetDetectedData(SimFusion_T* fusion);
	SIM_API int SimFusion_GetDetectedLength(SimFusion_T* fusion);


#ifdef __cplusplus
}
#endif
