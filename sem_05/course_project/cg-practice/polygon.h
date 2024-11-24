#ifndef POLYGON_H
#define POLYGON_H

#include <QVector3D>
#include <cstdlib>
#include <tuple>
#include <vector>

#include "lightsource.h"
#include "sight.h"

class Polygon {
public:
    Polygon();
    Polygon(std::vector<size_t> vertices, std::vector<QVector3D> points);
    Polygon(const Polygon& polygon);
    Polygon(Polygon&& polygon);

    Polygon& operator=(const Polygon& polygon);
    Polygon& operator=(Polygon&& polygon);

    std::vector<size_t> getVertices();

    std::tuple<bool, double, QVector3D> rayIntersection(const sight_t& sight, const std::vector<QVector3D>& points, const std::vector<LightSource>& lights, const double& specCoef);
    bool isInside(QVector3D& point, const std::vector<QVector3D>& points);
    void norm(std::vector<QVector3D>& points);

private:
    std::vector<size_t> vertices;
    double a = 0, b = 0, c = 0, d = 0;
};

#endif // POLYGON_H
