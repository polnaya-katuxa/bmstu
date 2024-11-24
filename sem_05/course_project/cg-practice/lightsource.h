#ifndef LIGHTSOURCE_H
#define LIGHTSOURCE_H

#include <QVector3D>

class LightSource {
public:
    LightSource();
    LightSource(const QVector3D& pos, const double& intensity)
        : pos(pos)
        , intensity(intensity)
    {
    }

    QVector3D getPos();
    double getIntensity();

private:
    QVector3D pos;
    double intensity;
};

#endif // LIGHTSOURCE_H
