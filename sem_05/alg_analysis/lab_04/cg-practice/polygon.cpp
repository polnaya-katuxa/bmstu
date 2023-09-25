#include "polygon.h"
#include "math.h"
#include <QtGlobal>

QVector3D calculateMiddlePoint(std::vector<QVector3D> pointsVec)
{
    QVector3D middlePoint = QVector3D(0, 0, 0);

    for (size_t i = 0; i < pointsVec.size(); i++) {
        middlePoint = middlePoint + pointsVec[i];
    }
    middlePoint.setX(middlePoint.x() / pointsVec.size());
    middlePoint.setY(middlePoint.y() / pointsVec.size());
    middlePoint.setZ(middlePoint.z() / pointsVec.size());

    return middlePoint;
}

Polygon::Polygon()
{
}

Polygon::Polygon(std::vector<size_t> vertices, std::vector<QVector3D> points)
    : vertices(vertices)
{
    QVector3D mp = calculateMiddlePoint(points);
    QVector3D p0 = points[vertices[0]];
    QVector3D p1 = points[vertices[1]];
    QVector3D p2 = points[vertices[2]];

    double kx = (p1.y() - p0.y()) * (p2.z() - p0.z()) - (p2.y() - p0.y()) * (p1.z() - p0.z());
    double ky = (p1.x() - p0.x()) * (p2.z() - p0.z()) - (p2.x() - p0.x()) * (p1.z() - p0.z());
    double kz = (p1.x() - p0.x()) * (p2.y() - p0.y()) - (p2.x() - p0.x()) * (p1.y() - p0.y());

    a = kx;
    b = -ky;
    c = kz;
    d = -kx * p0.x() + ky * p0.y() - kz * p0.z();

    if (a * mp.x() + b * mp.y() + c * mp.z() + d > 0.0) {
        a = -a;
        b = -b;
        c = -c;
        d = -d;
    }
}

Polygon::Polygon(const Polygon& polygon)
{
    vertices = polygon.vertices;
    a = polygon.a;
    b = polygon.b;
    c = polygon.c;
    d = polygon.d;
}

Polygon::Polygon(Polygon&& polygon)
{
    vertices = polygon.vertices;
    polygon.vertices.clear();
    a = polygon.a;
    b = polygon.b;
    c = polygon.c;
    d = polygon.d;
}

Polygon& Polygon::operator=(const Polygon& polygon)
{
    vertices = polygon.vertices;
    a = polygon.a;
    b = polygon.b;
    c = polygon.c;
    d = polygon.d;
    return *this;
}

Polygon& Polygon::operator=(Polygon&& polygon)
{
    vertices = polygon.vertices;
    polygon.vertices.clear();
    a = polygon.a;
    b = polygon.b;
    c = polygon.c;
    d = polygon.d;
    return *this;
}

std::vector<size_t> Polygon::getVertices()
{
    return vertices;
}

inline int sign(int num)
{
    if (num < 0)
        return -1;

    if (num > 0)
        return 1;

    return 0;
}

bool Polygon::isInside(QVector3D& point, const std::vector<QVector3D>& points)
{
    int signX = 0;
    int signY = 0;
    int signZ = 0;

    for (size_t i = 0; i < vertices.size(); i++) {
        QVector3D p0 = points[vertices[i]];
        QVector3D p1 = points[vertices[(i + 1) % vertices.size()]];

        QVector3D side = p1 - p0;
        QVector3D toPoint = point - p0;

        QVector3D prod = QVector3D::crossProduct(side, toPoint);

        int curSignX = sign(prod.x());
        int curSignY = sign(prod.y());
        int curSignZ = sign(prod.z());

        if (curSignX != 0) {
            if (curSignX != signX && signX != 0)
                return false;
            signX = curSignX;
        }

        if (curSignY != 0) {
            if (curSignY != signY && signY != 0)
                return false;
            signY = curSignY;
        }

        if (curSignZ != 0) {
            if (curSignZ != signZ && signZ != 0)
                return false;
            signZ = curSignZ;
        }
    }

    return true;
}

void Polygon::norm(std::vector<QVector3D>& points)
{
    QVector3D mp = calculateMiddlePoint(points);
    QVector3D p0 = points[vertices[0]];
    QVector3D p1 = points[vertices[1]];
    QVector3D p2 = points[vertices[2]];

    double kx = (p1.y() - p0.y()) * (p2.z() - p0.z()) - (p2.y() - p0.y()) * (p1.z() - p0.z());
    double ky = (p1.x() - p0.x()) * (p2.z() - p0.z()) - (p2.x() - p0.x()) * (p1.z() - p0.z());
    double kz = (p1.x() - p0.x()) * (p2.y() - p0.y()) - (p2.x() - p0.x()) * (p1.y() - p0.y());

    a = kx;
    b = -ky;
    c = kz;
    d = -kx * p0.x() + ky * p0.y() - kz * p0.z();

    if (a * mp.x() + b * mp.y() + c * mp.z() + d > 0.0) {
        a = -a;
        b = -b;
        c = -c;
        d = -d;
    }
}

std::tuple<bool, double, QVector3D> Polygon::rayIntersection(const sight_t& sight, const std::vector<QVector3D>& points, const std::vector<LightSource>& lights, const double& specCoef)
{
    QVector3D dir = sight.dir;
    QVector3D start = sight.cam;
    dir = dir.normalized();

    if (vertices.size() < 3)
        return std::tuple<bool, double, QVector3D>(false, 0.0, QVector3D(0, 0, 0));

    double k = a * dir.x() + b * dir.y() + c * dir.z();
    if (qFuzzyIsNull(k))
        return std::tuple<bool, double, QVector3D>(false, 0.0, QVector3D(0, 0, 0));

    double t = -(a * start.x() + b * start.y() + c * start.z() + d) / k;
    if (t < 0)
        return std::tuple<bool, double, QVector3D>(false, 0.0, QVector3D(0, 0, 0));

    QVector3D intersection = start + dir * t;
    if (!isInside(intersection, points))
        return std::tuple<bool, double, QVector3D>(false, 0.0, QVector3D(0, 0, 0));

    QVector3D norm = QVector3D(a, b, c).normalized();

    return std::tuple<bool, double, QVector3D>(true, t, norm);
}
