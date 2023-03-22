#pragma once

#include "SimExport.h"
#include "SimData_cgo.h"

typedef struct SimForce_T SimForce_T;

#ifdef __cplusplus
extern "C" {
#endif

	SIM_API SimForce_T* SimForce_New();
	SIM_API unsigned int SimForce_Id(SimForce_T* force);
	SIM_API unsigned int SimForce_PlatformCode(SimForce_T* force);
	SIM_API const char* SimForce_PlatformType(SimForce_T* force);
	SIM_API const char* SimForce_Name(SimForce_T* force);
	SIM_API const char* SimForce_Team(SimForce_T* force);
	SIM_API double SimForce_Lon(SimForce_T* force);
	SIM_API double SimForce_Lat(SimForce_T* force);
	SIM_API double SimForce_Alt(SimForce_T* force);
	SIM_API double SimForce_Heading(SimForce_T* force);
	SIM_API double SimForce_Pitch(SimForce_T* force);
	SIM_API double SimForce_Roll(SimForce_T* force);
	SIM_API double SimForce_Speed(SimForce_T* force);
	SIM_API double SimForce_Life(SimForce_T* force);
	SIM_API void SimForce_SetId(SimForce_T* force, unsigned int id);
	SIM_API void SimForce_SetPlatformCode(SimForce_T* force, unsigned int code);
	SIM_API void SimForce_SetPlatformType(SimForce_T* force, const char* type);
	SIM_API void SimForce_SetName(SimForce_T* force, const char* name);
	SIM_API void SimForce_SetTeam(SimForce_T* force, const char* team);
	SIM_API void SimForce_SetLon(SimForce_T* force, double lon);
	SIM_API void SimForce_SetLat(SimForce_T* force, double lat);
	SIM_API void SimForce_SetAlt(SimForce_T* force, double alt);
	SIM_API void SimForce_SetHeading(SimForce_T* force, double heading);
	SIM_API void SimForce_SetPitch(SimForce_T* force, double pitch);
	SIM_API void SimForce_SetRoll(SimForce_T* force, double roll);
	SIM_API void SimForce_SetSpeed(SimForce_T* force, double speed);
	SIM_API void SimForce_SetLife(SimForce_T* force, double life);
	SIM_API void SimForce_NextStep(SimForce_T* force, long long simTime, long long step);
	SIM_API void SimForce_DoAction(SimForce_T* force, const char* action, const char* params);
	SIM_API void SimForce_Prepared(SimForce_T* force);

	SIM_API void SimForce_SetMaxDRAerial(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxDRGround(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxDRSea(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxDRUnderSea(SimForce_T* force, double value);
	SIM_API void SimForce_SetRateDeflect(SimForce_T* force, double value);
	SIM_API void SimForce_SetRatePitch(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxDeflect(SimForce_T* force, double value);
	SIM_API void SimForce_SetMinDeflect(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxPitch(SimForce_T* force, double value);
	SIM_API void SimForce_SetMinPitch(SimForce_T* force, double value);
	SIM_API void SimForce_SetPrecisionPosition(SimForce_T* force, double value);
	SIM_API void SimForce_SetPrecisionVelocity(SimForce_T* force, double value);
	SIM_API void SimForce_SetPrecisionDirect(SimForce_T* force, double value);
	SIM_API void SimForce_SetResolution(SimForce_T* force, double value);

	SIM_API void SimForce_SetMaxSRAerial(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxSRGround(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxSRSea(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxSRUnderSea(SimForce_T* force, double value);

	SIM_API void SimForce_SetMaxRollRate(SimForce_T* force, double value);
	SIM_API void SimForce_SetTurnRadiusHori(SimForce_T* force, double value);
	SIM_API void SimForce_SetTurnRadiusVert(SimForce_T* force, double value);
	SIM_API void SimForce_SetTurnRateHori(SimForce_T* force, double value);
	SIM_API void SimForce_SetTurnRateVert(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxSpeedUp(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxSpeedDown(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxSpeed(SimForce_T* force, double value);
	SIM_API void SimForce_SetPatrolSpeed(SimForce_T* force, double value);
	SIM_API void SimForce_SetMinSpeed(SimForce_T* force, double value);
	SIM_API void SimForce_SetMaxMileage(SimForce_T* force, double value);
	SIM_API void SimForce_SetManeuvering(SimForce_T* force, int value);

	SIM_API void SimForce_SetPosture(SimForce_T* force, const ForcePosture params);
	SIM_API ForcePosture SimForce_Posture(SimForce_T* force);

#ifdef __cplusplus
}
#endif