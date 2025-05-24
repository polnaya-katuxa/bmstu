# Лабораторная работа 3. Вариант 7.

function main()
  clc;

  debug = true;

  a = 1;
  b = 2;
  eps = 1e-6;

  [x_min, f_min, n, xs, fs, as, bs] = find_min(debug, a, b, eps);
  draw_plot(debug, a, b, eps, x_min, f_min, xs, fs, as, bs);
  fprintf('\n\033[36mТочка минимума (x*, f(x*)) = (%.10f, %.10f), количество вычислений функции: %d.\033[0m\n', x_min, f_min, n);
end

function [x1, x2, x3, f1, f2, f3, n] = find_x123(debug, a, b)
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

  fa = f(a);
  fb = f(b);

  while true
    if f1 <= f2
      b = x2;
      fb = f2;
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
      f1 = f1;
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

    if (fa > f1 && f1 <= f2) || (fa >= f1 && f1 < f2)
      x3 = x2;
      f3 = f2;

      x2 = x1;
      f2 = f1;

      x1 = a;
      f1 = fa;

      n = i;
      return;
    elseif (f1 > f2 && f2 <= fb) || (f1 >= f2 && f2 < fb)
      x3 = b;
      f3 = fb;

      n = i;
      return;
    endif
  endwhile
end

function [x_min, f_min, n, xs, fs, as, bs] = find_min(debug, a, b, eps)
  [x1, x2, x3, f1, f2, f3, n] = find_x123(debug, a, b);

  xs = [];
  fs = [];
  as = [];
  bs = [];

  as(end + 1) = x1;
  bs(end + 1) = x3;

  a1 = (f2 - f1) / (x2 - x1);
  a2 = 1 / (x3 - x2) * ((f3 - f1) / (x3 - x1) - (f2 - f1) / (x2 - x1));
  x_prob = 1/2 * (x1 + x2 - a1/a2);
  f_prob = f(x_prob);
  n = n + 1;

  if debug
    fprintf('Итерация %d: (x1, f1) = (%.10f, %.10f), (x2, f2) = (%.10f, %.10f), (x3, f3) = (%.10f, %.10f), (x_prob, f_prob) = (%.10f, %.10f).\n', n-1, x1, f1, x2, f2, x3, f3, x_prob, f_prob);
  endif

  xs(end + 1) = x_prob;
  fs(end + 1) = f_prob;

  while true
    x_old = x_prob;

    if x2 < x_prob
      if f2 < f_prob
        x3 = x_prob;
        f3 = f_prob;
      else
        x1 = x2;
        f1 = f2;
        x2 = x_prob;
        f2 = f_prob;
      endif
    else
      if f_prob < f2
        x3 = x2;
        f3 = f2;
        x2 = x_prob;
        f2 = f_prob;
      else
        x1 = x_prob;
        f1 = f_prob;
      endif
    endif

    as(end + 1) = x1;
    bs(end + 1) = x3;

    a1 = (f2 - f1) / (x2 - x1);
    a2 = 1 / (x3 - x2) * ((f3 - f1) / (x3 - x1) - (f2 - f1) / (x2 - x1));
    x_prob = 1/2 * (x1 + x2 - a1/a2);
    f_prob = f(x_prob);
    n = n + 1;
    if debug
      fprintf('Итерация %d: (x1, f1) = (%.10f, %.10f), (x2, f2) = (%.10f, %.10f), (x3, f3) = (%.10f, %.10f), (x_prob, f_prob) = (%.10f, %.10f).\n', n-1, x1, f1, x2, f2, x3, f3, x_prob, f_prob);
    endif

    xs(end + 1) = x_prob;
    fs(end + 1) = f_prob;

    if abs(x_prob - x_old) < eps
      x_min = x_prob;
      f_min = f_prob;
      return;
    endif
  endwhile
end

function draw_plot(debug, a, b, step, x_min, f_min, xs, fs, as, bs)
  sleep_seconds = 7;

  %fprintf('as %d, xs %d\n', length(as), length(xs));

  x = a:step:b;
  y = zeros(size(x));
  for i = 1:length(x)
      y(i) = f(x(i));
  end
  plot(x,y);
  hold on;

  if debug
    line([as(1), bs(1)], [f(as(1)), f(bs(1))], 'DisplayName', sprintf('шаг %d', 1), 'Color', 'r');
    scatter(xs(1), fs(1), 8, 'r', 'filled');
    pause(sleep_seconds);

    for i = 2:length(xs)
        if i < length(as)
          line([as(i), bs(i)], [f(as(i)), f(bs(i))], 'DisplayName', sprintf('шаг %d', i), 'Color', 'r');
        endif

        pause(sleep_seconds);

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
  y = atan(x .^ 3 - 5 * x + 1) + ((x .^ 2) / (3 * x - 2)) .^ sqrt(3);
end



