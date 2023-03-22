#pragma once

// ������̬���ݣ�ÿ�������ڵ�ͬ��
typedef struct ForcePosture
{
	unsigned int id;			// ����id
	double lon;					// ����
	double lat;					// γ��
	double alt;					// �߶�
	double heading;				// �����
	double pitch;				// ƫת��
	double roll;				// ��ת��
	double speed;				// �ٶ�
	double life;				// ����ֵ
	double remainingMileage;	// ʣ���������
} ForcePosture;

// ̽�⵽�ı���̬��
typedef struct DetectedPostore
{
	unsigned int id;			// ����id
	double lon;					// ����
	double lat;					// γ��
	double alt;					// �߶�
	double heading;				// �����
	double pitch;				// ƫת��
	double speed;				// �ٶ�
} DetectedPostore;
