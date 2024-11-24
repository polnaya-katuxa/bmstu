x = [7.76,6.34,5.11,7.62,8.84,4.68,8.65,6.90,8.79,6.61,6.62,7.13,6.75,...
    7.28,7.74,7.08,5.57,8.20,7.78,7.92,6.00,4.88,6.75,6.56,7.48,8.51,...
    9.06,6.94,6.93,7.79,5.71,5.93,6.81,5.76,5.88,7.05,7.22,6.67,5.59,...
    6.57,7.28,6.22,6.31,5.51,6.69,7.12,7.40,6.86,7.28,6.82,7.08,7.52,...
    6.81,7.55,4.89,5.48,7.74,5.10,8.17,7.67,7.07,5.80,6.10,7.15,7.88,...
    9.06,6.85,4.88,6.74,8.76,8.53,6.72,7.21,7.42,8.29,8.56,9.25,6.63,...
    7.49,6.67,6.79,5.19,8.20,7.97,8.64,7.36,6.72,5.90,5.53,6.44,7.35,...
    5.18,8.25,5.68,6.29,6.69,6.08,7.42,7.10,7.14,7.10,6.60,6.35,5.99,...
    6.17,9.05,6.01,7.77,6.27,5.81,7.80,9.89,4.39,6.83,6.53,8.15,6.68,...
    6.87,6.31,6.83];
%x = 3*x;

prompt = "Enter gamma: ";
gamma = input(prompt)

n = length(x);

mu = sum(x) / n
s2 = sum((x - mu) .^ 2) / (n - 1)

mu_low = mu - sqrt(s2) * tinv((1 + gamma) / 2, n - 1) / sqrt(n)
mu_high = mu + sqrt(s2) * tinv((1 + gamma) / 2, n - 1) / sqrt(n)

s2_low = (n - 1) * s2 / chi2inv((1 + gamma) / 2, n - 1)
s2_high = (n - 1) * s2 / chi2inv((1 - gamma) / 2, n - 1)

n_arr = zeros([1 n]);
mu_arr = zeros([1 n]);
mu_low_arr = zeros([1 n]);
mu_high_arr = zeros([1 n]);
s2_arr = zeros([1 n]);
s2_low_arr = zeros([1 n]);
s2_high_arr = zeros([1 n]);
mu_narr = zeros([1 n]);
s2_narr = zeros([1 n]);

for i = 1:n
    n_arr(i) = i;

    mu_narr(i) = mu;
    s2_narr(i) = s2;

    mu_arr(i) = sum(x(1:i)) / i;
    s2_arr(i) = sum((x(1:i) - mu) .^ 2) / (i - 1);

    mu_low_arr(i) = mu_arr(i) - sqrt(s2_arr(i)) * ...
        tinv((1 + gamma) / 2, i - 1) / sqrt(i);
    mu_high_arr(i) = mu_arr(i) + sqrt(s2_arr(i)) * ...
        tinv((1 + gamma) / 2, i - 1) / sqrt(i);

    s2_low_arr(i) = (i - 1) * s2_arr(i) / ...
        chi2inv((1 + gamma) / 2, i - 1);
    s2_high_arr(i) = (i - 1) * s2_arr(i) / ...
        chi2inv((1 - gamma) / 2, i - 1);
end

start = 15;

plot(start:n, mu_narr(start:n), start:n, mu_arr(start:n), ...
    start:n, mu_low_arr(start:n), start:n, mu_high_arr(start:n), 'LineWidth', 1.5);
title('Graphic for $\hat \mu$','interpreter','latex', 'FontSize', 14);
xlabel('n');
ylabel('y');
xlim([15 n]);
legend('$\hat \mu(\vec x_N)$', '$\hat \mu(\vec x_n)$', ...
    '$\underline{\mu}(\vec x_n)$', '$\overline{\mu}(\vec x_n)$', ...
    'Interpreter', 'latex', 'FontSize', 14);
figure;

plot(start:n, s2_narr(start:n), start:n, s2_arr(start:n), ...
    start:n, s2_low_arr(start:n), start:n, s2_high_arr(start:n), 'LineWidth', 1.5);
title('Graphic for $\hat S^2$','interpreter','latex', 'FontSize', 14);
xlabel('n');
ylabel('z');
xlim([15 n]);
legend('$\hat S^2(\vec x_N)$', '$\hat S^2(\vec x_n)$', ...
   '$\underline{\sigma}^2(\vec x_n)$', ...
   '$\overline{\sigma}^2(\vec x_n)$', ...
   'Interpreter', 'latex', 'FontSize', 14);




