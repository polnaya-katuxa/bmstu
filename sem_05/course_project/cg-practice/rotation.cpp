#include "rotation.h"
#include <iostream>
#include <math.h>

inline double toRad(double angle)
{
    return angle * M_PI / 180.0;
}

void rotate(double& x, double& y, const double& angle)
{
    double angleRads = toRad(angle);
    double tempX = x;
    double tempY = y;

    x = tempX * cos(angleRads) + tempY * sin(angleRads);
    y = -tempX * sin(angleRads) + tempY * cos(angleRads);
}

void pointRotate(QVector3D& point, QVector3D angles)
{
    double x = point.x();
    double y = point.y();
    double z = point.z();

    rotate(x, y, angles.z());
    rotate(z, x, angles.y());
    rotate(y, z, angles.x());

    point.setX(x);
    point.setY(y);
    point.setZ(z);
}
