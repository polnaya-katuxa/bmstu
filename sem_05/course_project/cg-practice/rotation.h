#ifndef ROTATION_H
#define ROTATION_H

#include <QVector3D>

void rotate(double& x, double& y, const double& angle);

void pointRotate(QVector3D& point, QVector3D angles);

#endif // ROTATION_H
