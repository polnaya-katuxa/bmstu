#include "sphere.h"

#include <limits>
#include <math.h>

Sphere::Sphere()
{
}

Sphere::Sphere(QVector3D cen, double rad)
    : cen(cen)
    , rad(rad)
{
}

std::tuple<bool, double, QVector3D> Sphere::rayIntersection(const sight_t& sight, const std::vector<LightSource>& lights)
{
    double t = 0.0;
    QVector3D start = sight.cam;
    QVector3D dir = sight.dir;
    QVector3D norm = QVector3D(0, 0, 0);
    QVector3D pt = QVector3D(0, 0, 0);

    QVector3D L = cen - start; // соед. центр и начало луча
    double tca = QVector3D::dotProduct(L, dir); // проекция L на луч (скаляр.произв.)
    double d2 = QVector3D::dotProduct(L, L) - tca * tca; // ближайшая к центру сферы точка луча (ее сдвиг от начала,
    // как сторона прямоуг. треугольника)

    double rad2 = rad * rad;
    if (d2 > rad2) // пересечение вне сферы, не пересекает типа
        return std::tuple<bool, double, QVector3D>(false, t, norm);

    if (rad2 - d2 < 0)
        return std::tuple<bool, double, QVector3D>(false, t, norm);
    double thc = sqrt(rad2 - d2); // расстояние от той ближайшей точки до пересечения

    double t0 = tca - thc, t1 = tca + thc; // 0 - если внутри сферы, 1 - если снаружи
    if (t0 > 1.0)
        t = t0;
    else if (t1 > 1.0)
        t = t1;
    else
        return std::tuple<bool, double, QVector3D>(false, t, norm);

    pt = start + dir * t;
    norm = (pt + (cen * (-1))).normalized();

    return std::tuple<bool, double, QVector3D>(true, t, norm);
}

void Sphere::rotate(QVector3D& angles)
{
    pointRotate(cen, angles);
}

QVector3D Sphere::getCen()
{
    return cen;
}

double Sphere::getRad()
{
    return rad;
}
