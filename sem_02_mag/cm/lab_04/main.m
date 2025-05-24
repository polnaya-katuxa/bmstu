# Лабораторная работа 4. Вариант 7.

function main()
  clc;

  debug = true;

  sleep_seconds = 2;
  a = 1;
  b = 2;
  eps = 1e-6;

  [x_min, f_min, n, xs, fs] = find_min(debug, a, b, eps);
  fprintf('\n\033[36mТочка минимума (x*, f(x*)) = (%.10f, %.10f), количество вычислений функции: %d.\033[0m\n', x_min, f_min, n);

  options = optimset('TolX', eps);
  if debug
    options = optimset(options, 'Display', 'iter');
  end
  [x_min_default, f_min_default] = fminbnd(@f, a, b, options);
  fprintf('\n\033[36mТочка минимума методом fminbnd (x*, f(x*)) = (%.10f, %.10f).\033[0m\n', x_min_default, f_min_default);

  draw_plot(debug, a, b, eps, x_min, f_min, xs, fs, sleep_seconds);

  x_min + 0.777
end

function [x_min, f_min, n, xs, fs] = find_min(debug, a, b, eps)
  [x_prob, f_prob, n] = find_x0(debug, a, b);

  #x_prob = (a + b) / 2;
  #f_prob = f(x_prob);
  #n = 1;

  xs = [];
  fs = [];

  xs(end + 1) = x_prob;
  fs(end + 1) = f_prob;

  delta = 1e-3;

  i = n - 1;
  if debug
    fprintf('Итерация %d: (x, f) = (%.10f, %.10f).\n', i, x_prob, f_prob);
  endif
  i = i + 1;

  f2 = (f(x_prob - delta) - 2 * f_prob + f(x_prob + delta)) / (delta ^ 2);
  n = n + 2;

  while true
    f1 = (f(x_prob + delta) - f(x_prob - delta)) / (2 * delta)
    n = n + 2;

    if abs(f1) < eps
      x_min = x_prob;
      f_min = f(x_prob);
      n = n + 1;
      return;
    endif

    x_prob = x_prob - f1/f2;
    xs(end + 1) = x_prob;
    fs(end + 1) = f(x_prob);
    if debug
      fprintf('Итерация %d: (x, f) = (%.10f, %.10f).\n', i, x_prob, f(x_prob));
    endif
    i = i+1;
  endwhile
end

function draw_plot(debug, a, b, step, x_min, f_min, xs, fs, sleep_seconds)
  x=a:step:b;
  y = zeros(size(x));
  for i = 1:length(x)
      y(i) = f(x(i));
  end
  plot(x,y);
  hold on;
  if debug
    scatter(xs(1), fs(1), 8, 'r', 'filled');
    for i = 2:length(xs)
        scatter(xs(i-1), fs(i-1), 8, 'b', 'filled');
        scatter(xs(i), fs(i), 8, 'r', 'filled');
        pause(sleep_seconds);
    end
  endif
  scatter(x_min, f_min, 10, 'g', 'filled');
  text(x_min, f_min, sprintf('\n\n\n\n(%.10f, %.10f)', x_min, f_min), 'FontSize', 12);
  hold off;
end

function y = f(x)
  %y = atan(x .^ 3 - 5 * x + 1) + ((x .^ 2) / (3 * x - 2)) .^ sqrt(3);
  y = (x+0.777).^4;
end

function [x0, f0, n] = find_x0(debug, a, b)
  iterations = 4;

  tau = (sqrt(5) - 1) / 2;

  l = b - a;

  x1 = b - tau*l;
  f1 = f(x1);

  if debug
    fprintf('Золотое сечение (x0, f(x0)) = (%f, %f).\n', x1, f1);
  endif

  x2 = a + tau*l;
  f2 = f(x2);

  if debug
    fprintf('Золотое сечение (x1, f(x1)) = (%f, %f).\n', x2, f2);
  endif

  i = 2;

  for j = 1:iterations
    if f1 <= f2
      b = x2;
      l = b - a;

      x2 = x1;
      f2 = f1;

      x1 = b - tau*l;
      f1 = f(x1);
      i = i + 1;

      if debug
        fprintf('Золотое сечение (x%d, f(x%d)) = (%f, %f).\n', i-1, i-1, x1, f1);
      endif
    else
      a = x1;
      l = b - a;

      x1 = x2;
      f1 = f2;

      x2 = a + tau*l;
      f2 = f(x2);
      i = i + 1;

      if debug
        fprintf('Золотое сечение (x%d, f(x%d)) = (%f, %f).\n', i-1, i-1, x2, f2);
      endif
    endif
  endfor

  n = i + 1;
  x0 = (a + b) / 2;
  f0 = f(x0);
end

