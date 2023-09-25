QT       += core gui testlib

greaterThan(QT_MAJOR_VERSION, 4): QT += widgets

CONFIG += c++17

# You can make your code fail to compile if it uses deprecated APIs.
# In order to do so, uncomment the following line.
#DEFINES += QT_DISABLE_DEPRECATED_BEFORE=0x060000    # disables all the APIs deprecated before Qt 6.0.0

SOURCES += \
    drawing.cpp \
    figure.cpp \
    lightsource.cpp \
    main.cpp \
    mainwindow.cpp \
    material.cpp \
    polygon.cpp \
    polyhedron.cpp \
    rotation.cpp \
    sphere.cpp \
    test_rotation.cpp

HEADERS += \
    drawing.h \
    figure.h \
    lightsource.h \
    mainwindow.h \
    material.h \
    objloader.h \
    polygon.h \
    polyhedron.h \
    rotation.h \
    sight.h \
    sphere.h \
    test_rotation.h

FORMS += \
    mainwindow.ui

# Default rules for deployment.
qnx: target.path = /tmp/$${TARGET}/bin
else: unix:!android: target.path = /opt/$${TARGET}/bin
!isEmpty(target.path): INSTALLS += target
