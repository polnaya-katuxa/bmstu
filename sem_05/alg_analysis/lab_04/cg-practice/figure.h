#ifndef FIGURE_H
#define FIGURE_H

#include "material.h"
#include "polygon.h"

#include <QColor>
#include <QPoint>
#include <QVector3D>
#include <QVector4D>
#include <vector>

class Figure {
public:
    Figure();

    Material getMaterial();
    void setMaterial(QVector4D alb, QColor dC, QColor sC, double rCoef, double sCoef);

    virtual std::tuple<bool, double, QVector3D> rayIntersection(const sight_t& sight, const std::vector<LightSource>& lights) = 0;
    virtual void rotate(QVector3D& angles) = 0;

protected:
    Material material;
};

#endif // FIGURE_H
