#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include "drawing.h"
#include <QMainWindow>

QT_BEGIN_NAMESPACE
namespace Ui {
class MainWindow;
}
QT_END_NAMESPACE

class MainWindow : public QMainWindow {
    Q_OBJECT

public:
    MainWindow(QWidget* parent = nullptr);
    ~MainWindow();

    void showEvent(QShowEvent* ev);

private slots:
    void on_drawButton_clicked();

    void on_actionSave_triggered();

    void on_actionLoad_triggered();

    void windowShown();

private:
    Ui::MainWindow* ui;

    Drawing picture;

    std::shared_ptr<QImage> curImage = nullptr;
};
#endif // MAINWINDOW_H
