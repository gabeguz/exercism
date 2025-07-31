defmodule FreelancerRates do
  def daily_rate(hourly_rate) do
    8.0 * hourly_rate
  end

  def apply_discount(before_discount, discount) do
    before_discount - before_discount * (discount / 100)
  end

  def monthly_rate(hourly_rate, discount) do
    total = 22 * daily_rate(hourly_rate)
    if discount > 0 do
       trunc(ceil(apply_discount(total, discount)))
     else 
       trunc(ceil(total))
     end
  end

  def days_in_budget(budget, hourly_rate, discount) do
    if discount > 0 do 
      Float.floor(budget / apply_discount(daily_rate(hourly_rate), discount), 1)
    else 
      Float.floor(budget / daily_rate(hourly_rate), 1)  
    end
  end
end