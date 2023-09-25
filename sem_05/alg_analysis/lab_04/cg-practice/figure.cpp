#include "figure.h"
#include <limits>

Figure::Figure()
{
}

Material Figure::getMaterial()
{
    return material;
}

void Figure::setMaterial(QVector4D alb, QColor dC, QColor sC, double rCoef, double sCoef)
{
    material.setAlbedo(alb);
    material.setDiffColor(dC);
    material.setSpecColor(sC);
    material.setRefCoef(rCoef);
    material.setSpecCoef(sCoef);
}
