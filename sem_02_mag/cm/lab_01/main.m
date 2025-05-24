# Лабораторная работа 1. Вариант 7.

function main()
  clc;

  debug = false;

  a = 1;
  b = 2;
  eps = 1e-6;

  [x_min, f_min, n, xs, fs] = find_min(debug, a, b, eps);
  draw_plot(a, b, eps, x_min, f_min, xs, fs);

  fprintf('\n\033[36mТочка минимума (x*, f(x*)) = (%.10f, %.10f), количество вычислений функции: %d.\033[0m\n', x_min, f_min, n);
  %x_min - 1.777
  end

function [x_min, f_min, n, xs, fs] = find_min(debug, a, b, eps)
  x0 = a;
  f0 = f(x0);
  delta = (b - a)/4;

  xs = [];
  fs = [];
  xs(end + 1) = x0;
  fs(end + 1) = f0;

  if debug
    fprintf('(x0, f(x0)) = (%.10f, %.10f).\n', x0, f0);
  endif

  i = 1;

  while true
    x1 = x0 + delta;
    f1 = f(x1);
    i = i + 1;

    if debug
      fprintf('(x%d, f(x%d)) = (%.10f, %.10f).\n', i-1, i-1, x1, f1);
    endif

    if f0 > f1
      if a <= x1 <= b
        x0 = x1;
        f0 = f1;

        xs(end + 1) = x1;
        fs(end + 1) = f1;

        continue;
      endif
    endif

    if abs(delta) < eps
      x_min = x0;
      f_min = f0;
      n = i;
      return;
    endif

    delta = -delta / 4;
    x0 = x1;
    f0 = f1;

    xs(end + 1) = x1;
    fs(end + 1) = f1;
  endwhile
end

function draw_plot(a, b, step, x_min, f_min, xs, fs)
  x=a:step:b;
  y = zeros(size(x));
  for i = 1:length(x)
      y(i) = f(x(i));
  end
  plot(x,y);
  hold on;
  scatter(xs(1), fs(1), 8, 'r', 'filled');
  for i = 2:length(xs)
      scatter(xs(i-1), fs(i-1), 8, 'b', 'filled');
      scatter(xs(i), fs(i), 8, 'r', 'filled');
      pause(2);
  end
  scatter(x_min, f_min, 10, 'g', 'filled');
  text(x_min, f_min, sprintf('\n\n\n\n(%.3f, %.3f)', x_min, f_min), 'FontSize', 12);
  hold off;
end

function y = f(x)
  y = atan(x .^ 3 - 5 * x + 1) + ((x .^ 2) / (3 * x - 2)) .^ sqrt(3);
  %y = (x-1.777).^2;
end




