#include "material.h"

Material::Material()
{
}

Material::Material(QColor dC, QColor sC, QVector4D alb, double rCoef, double sCoef)
    : diffColor(dC)
    , specColor(sC)
    , albedo(alb)
    , refCoef(rCoef)
    , specCoef(sCoef)
{
}

Material::Material(const Material& material)
    : diffColor(material.diffColor)
    , specColor(material.specColor)
    , albedo(material.albedo)
    , refCoef(material.refCoef)
    , specCoef(material.specCoef)
{
}

Material& Material::operator=(const Material& material)
{
    diffColor = material.diffColor;
    specColor = material.specColor;
    albedo = material.albedo;
    refCoef = material.refCoef;
    specCoef = material.specCoef;
    return *this;
}

Material& Material::operator=(Material&& material)
{
    diffColor = material.diffColor;
    specColor = material.specColor;
    albedo = material.albedo;
    refCoef = material.refCoef;
    specCoef = material.specCoef;
    return *this;
}

QColor Material::getDiffColor()
{
    return diffColor;
}

QColor Material::getSpecColor()
{
    return specColor;
}

QVector4D Material::getAlbedo()
{
    return albedo;
}

double Material::getSpecCoef()
{
    return specCoef;
}

double Material::getRefCoef()
{
    return refCoef;
}

void Material::setDiffColor(QColor dC)
{
    diffColor = dC;
}

void Material::setSpecColor(QColor sC)
{
    specColor = sC;
}

void Material::setAlbedo(QVector4D a)
{
    albedo = a;
}

void Material::setSpecCoef(double sC)
{
    specCoef = sC;
}

void Material::setRefCoef(double rC)
{
    refCoef = rC;
}
