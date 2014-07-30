# 1st line
s = gets
s = s.chomp.split(" ")
param_num = s[0].to_i
user_num = s[1].to_i
top_num = s[2].to_i

# 2nd line
s = gets
s = s.chomp.split(" ")
item_point_arr = s

# data line
user_score = []
user_num.times{
  user_point = 0
  s = gets
  s = s.chomp.split(" ")
  param_num.times{|i|
    user_point += s[i].to_f * item_point_arr[i].to_f
  }
  user_score.push(user_point)
}

user_score.sort!{|a,b| b <=> a}
top_num.times {
  us = user_score.shift
  puts us.round
}
