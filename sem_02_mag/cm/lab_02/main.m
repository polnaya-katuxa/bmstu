# Лабораторная работа 2. Вариант 7.

function main()
  clc;

  debug = false;

  a = 1;
  b = 2;
  eps = 1e-6;

  [x_min, f_min, n, xs, fs, as, bs] = find_min(debug, a, b, eps);
  draw_plot(a, b, eps, x_min, f_min, xs, fs, as, bs);
  fprintf('\n\033[36mТочка минимума (x*, f(x*)) = (%.10f, %.10f), количество вычислений функции: %d.\033[0m\n', x_min, f_min, n);
  %x_min - 1.777
  end

function [x_min, f_min, n, xs, fs, as, bs] = find_min(debug, a, b, eps)
  tau = (sqrt(5) - 1) / 2;
  l = b - a;

  x1 = b - tau*l;
  f1 = f(x1);

  as = [];
  bs = [];
  as(end + 1) = a;
  bs(end + 1) = b;
  xs = [];
  fs = [];
  xs(end + 1) = x1;
  fs(end + 1) = f1;

  if debug
    fprintf('(x0, f(x0)) = (%f, %f).\n', x1, f1);
  endif

  x2 = a + tau*l;
  f2 = f(x2);
  xs(end + 1) = x2;
  fs(end + 1) = f2;

  if debug
    fprintf('(x1, f(x1)) = (%f, %f).\n', x2, f2);
  endif

  i = 2;

  while true
    if l <= 2 * eps
      x_min = (a + b) / 2;
      f_min = f(x_min);
      n = i + 1;
      return;
    endif

    if f1 <= f2
      b = x2;
      l = b - a;

      x2 = x1;
      f2 = f1;

      x1 = b - tau*l;
      f1 = f(x1);
      i = i + 1;

      xs(end + 1) = x1;
      fs(end + 1) = f1;
      as(end + 1) = a;
      bs(end + 1) = b;

      if debug
        fprintf('(x%d, f(x%d)) = (%f, %f).\n', i-1, i-1, x1, f1);
      endif
    else
      a = x1;
      l = b - a;

      x1 = x2;
      f1 = f2;

      x2 = a + tau*l;
      f2 = f(x2);
      i = i + 1;

      xs(end + 1) = x2;
      fs(end + 1) = f2;
      as(end + 1) = a;
      bs(end + 1) = b;

      if debug
        fprintf('(x%d, f(x%d)) = (%f, %f).\n', i-1, i-1, x2, f2);
      endif
    endif
  endwhile
end

function draw_plot(a, b, step, x_min, f_min, xs, fs, as, bs)
  fprintf('as %d, xs %d\n', length(as), length(xs));

  x = a:step:b;
  y = zeros(size(x));
  for i = 1:length(x)
      y(i) = f(x(i));
  end
  plot(x,y);
  hold on;

  scatter(xs(1), fs(1), 8, 'r', 'filled');
  line([as(1), bs(1)], [f(as(1)), f(bs(1))], 'DisplayName', sprintf('шаг %d', 1), 'Color', 'r');
  pause(0);

  for i = 2:length(xs)
      scatter(xs(i-1), fs(i-1), 8, 'b', 'filled');
      scatter(xs(i), fs(i), 8, 'r', 'filled');
      pause(0);

      if i < length(as)
        line([bs(i-1), as(i)], [f(bs(i-1)), f(as(i))], 'DisplayName', sprintf('шаг %d', i), 'Color', 'r');
        pause(0);
        line([as(i), bs(i)], [f(as(i)), f(bs(i))], 'DisplayName', sprintf('шаг %d', i), 'Color', 'r');
      endif

      pause(0);
  end

  scatter(x_min, f_min, 10, 'g', 'filled');
  text(x_min, f_min, sprintf('\n\n\n\n(%.10f, %.10f)', x_min, f_min), 'FontSize', 12);

  hold off;
end

function y = f(x)
  y = atan(x .^ 3 - 5 * x + 1) + ((x .^ 2) / (3 * x - 2)) .^ sqrt(3);
  %y = (x-1.777).^6;
end




