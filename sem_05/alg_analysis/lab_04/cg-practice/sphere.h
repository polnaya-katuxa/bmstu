#ifndef SPHERE_H
#define SPHERE_H

#include "figure.h"
#include "polygon.h"
#include "rotation.h"
#include "sight.h"
#include <QColor>
#include <QPoint>
#include <QVector3D>
#include <QVector4D>
#include <vector>

class Sphere : public Figure {
public:
    Sphere();
    Sphere(QVector3D cen, double rad);

    QVector3D getCen();
    double getRad();

    std::tuple<bool, double, QVector3D> rayIntersection(const sight_t& sight, const std::vector<LightSource>& lights);
    void rotate(QVector3D& angles);

private:
    QVector3D cen = QVector3D(0, 0, 0);
    double rad = 0.0;
};

#endif // SPHERE_H
