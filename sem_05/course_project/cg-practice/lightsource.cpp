#include "lightsource.h"

LightSource::LightSource()
{
}

QVector3D LightSource::getPos()
{
    return pos;
}

double LightSource::getIntensity()
{
    return intensity;
}
