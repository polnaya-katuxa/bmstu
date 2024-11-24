#include "polyhedron.h"
#include "material.h"
#include <limits>

Polyhedron::Polyhedron()
{
}

Polyhedron::Polyhedron(std::vector<QVector3D> pointsVec, std::vector<Polygon> polygonsVec)
    : points(pointsVec)
    , polygons(polygonsVec)
{
}

bool fuzzyIsZero(double n)
{
    return n > -1e-3 && n < 1e-3;
}

std::tuple<bool, double, QVector3D> Polyhedron::rayIntersection(const sight_t& sight, const std::vector<LightSource>& lights)
{
    bool flag = false;
    double t = std::numeric_limits<double>::max();
    QVector3D norm;

    for (size_t i = 0; i < polygons.size(); i++) {
        std::tuple<bool, double, QVector3D> res = polygons[i].rayIntersection(sight, points, lights, material.getSpecCoef());
        if (std::get<0>(res)) {
            double t1 = std::get<1>(res);
            if (t1 < t && !fuzzyIsZero(t1)) {
                t = t1;
                norm = std::get<2>(res);
                flag = true;
            }
        }
    }
    return std::tuple<bool, double, QVector3D>(flag, t, norm);
}

void Polyhedron::rotate(QVector3D& angles)
{
    for (auto& point : points) {
        pointRotate(point, angles);
    }
    for (auto& polygon : polygons) {
        polygon.norm(points);
    }
}

std::vector<QVector3D> Polyhedron::getPoints()
{
    return points;
}

std::vector<Polygon> Polyhedron::getPolygons()
{
    return polygons;
}

void Polyhedron::setPoints(std::vector<QVector3D> pointsVec)
{
    points = pointsVec;
}

void Polyhedron::setPolygons(std::vector<Polygon> polygonsVec)
{
    polygons = polygonsVec;
}
