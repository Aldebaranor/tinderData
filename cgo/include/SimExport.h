#pragma once

#if defined(_WIN32)
#if defined(SIM_EXPORT)
#define SIM_API __declspec(dllexport)
#else
#define SIM_API __declspec(dllimport)
#endif
#else
#define SIM_API __attribute__((visibility("default")))
#endif
