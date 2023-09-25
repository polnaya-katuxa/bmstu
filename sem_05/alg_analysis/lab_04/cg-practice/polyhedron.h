#ifndef POLYHEDRON_H
#define POLYHEDRON_H

#include "figure.h"
#include "polygon.h"
#include "rotation.h"
#include "sight.h"
#include <QColor>
#include <QPoint>
#include <QVector3D>
#include <QVector4D>
#include <vector>

class Polyhedron : public Figure {
public:
    Polyhedron();
    Polyhedron(std::vector<QVector3D> pointsVec, std::vector<Polygon> polygonsVec);

    std::vector<QVector3D> getPoints();
    std::vector<Polygon> getPolygons();

    void setPoints(std::vector<QVector3D> pointsVec);
    void setPolygons(std::vector<Polygon> polygonsVec);

    std::tuple<bool, double, QVector3D> rayIntersection(const sight_t& sight, const std::vector<LightSource>& lights);
    void rotate(QVector3D& angles);

private:
    std::vector<QVector3D> points;
    std::vector<Polygon> polygons;
};

#endif // POLYHEDRON_H
