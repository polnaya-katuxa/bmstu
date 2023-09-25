#include <QFileDialog>
#include <QMessageBox>
#include <QTimer>
#include <fstream>
#include <iostream>

#include "drawing.h"
#include "mainwindow.h"
#include "polyhedron.h"
#include "sphere.h"
#include "ui_mainwindow.h"

#define NUM_THREADS 8

MainWindow::MainWindow(QWidget* parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    QGraphicsScene* scene = new QGraphicsScene();
    ui->graphicsView->setScene(scene);
    ui->graphicsView->scene()->clear();
}

MainWindow::~MainWindow()
{
    delete ui->graphicsView->scene();
    delete ui;
}

void MainWindow::showEvent(QShowEvent* ev) // когда окно полностью сконструировано
{
    QMainWindow::showEvent(ev);
    QTimer::singleShot(50, this, SLOT(windowShown()));
}

void MainWindow::windowShown() // начальный вывод на экран
{
    picture = Drawing(ui->graphicsView->height(), ui->graphicsView->width());
    picture.initDefaultLightSources();
}

void MainWindow::on_drawButton_clicked()
{
    int x = ui->xBox->value();
    int y = ui->yBox->value();
    int z = ui->zBox->value();
    if (x || y || z) {
        picture.rotate(QVector3D(x, y, z));
    }
    curImage = picture.drawFigures(NUM_THREADS);
    QPixmap pixmap = QPixmap::fromImage(*curImage);
    ui->graphicsView->scene()->addPixmap(pixmap);
}

void MainWindow::on_actionSave_triggered()
{
    if (curImage == nullptr)
        ui->statusbar->showMessage("No image drawn. Cannot save white canvas.");
    else {
        QString filename = QFileDialog::getSaveFileName();
        if (!filename.isNull())
            curImage->save(filename);
        else
            ui->statusbar->showMessage("Problems with the file.");
    }
}

void MainWindow::on_actionLoad_triggered()
{
    QString filename = QFileDialog::getOpenFileName();

    if (!filename.isNull()) {
        if (!picture.loadFile(filename))
            QMessageBox::warning(this, "Error", "File loading error.");

    } else
        QMessageBox::warning(this, "Error", "Filename error.");
}
