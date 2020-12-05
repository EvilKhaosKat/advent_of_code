def solve_first
  passport = {}

  valid_passports = 0
  File.open('input.txt', 'r').each_line do |line|
    if line.length == 1
      valid_passports = valid_passports.next if valid_first passport
      passport = {}
    else
      line.split(' ').each do |pair|
        data = pair.split(':')
        passport[data[0]] = data[1]
      end
    end
  end
  valid_passports = valid_passports.next if valid_first passport

  valid_passports
end

def solve_second
  passport = {}

  valid_passports = 0
  File.open('input.txt', 'r').each_line do |line|
    if line.length == 1
      valid_passports = valid_passports.next if valid_second passport
      passport = {}
    else
      line.split(' ').each do |pair|
        data = pair.split(':')
        passport[data[0]] = data[1]
      end
    end
  end
  valid_passports = valid_passports.next if valid_second passport

  valid_passports
end

# @param [Dict] passport
def valid_first(passport)
  passport.key?('byr') \
     && passport.key?('iyr') \
     && passport.key?('eyr') \
     && passport.key?('hgt') \
     && passport.key?('hcl') \
     && passport.key?('ecl') \
     && passport.key?('pid')
end

def valid_second(passport)
  valid_byr(passport) \
     && valid_iyr(passport) \
     && valid_eyr(passport) \
     && valid_hgt(passport) \
     && valid_hcl(passport) \
     && valid_ecl(passport) \
     && valid_pid(passport)
end

def check_number(passport, key, length, min, max)
  value = passport[key]

  return false if value.nil?
  return false unless /\d{#{length}}/.match? value

  value = value.to_i
  value >= min && value <= max
end

def valid_byr(passport)
  check_number(passport, 'byr', 4, 1920, 2002)
end

def valid_iyr(passport)
  check_number(passport, 'iyr', 4, 2010, 2020)
end

def valid_eyr(passport)
  check_number(passport, 'eyr', 4, 2020, 2030)
end

def valid_hgt(passport)
  hgt_raw = passport['hgt']

  return false if hgt_raw.nil?
  return false unless /\d{2,3}(cm|in)/.match? hgt_raw

  value = hgt_raw[0...-2]
  type = hgt_raw[-2...hgt_raw.length]

  if type == 'cm'
    value.to_i.between?(150, 193)
  else
    value.to_i.between?(59, 76)
  end
end

def valid_hcl(passport)
  hcl = passport['hcl']

  return false if hcl.nil?

  /#[0-9a-f]{6}/.match? hcl
end

def valid_ecl(passport)
  ecl = passport['ecl']

  return false if ecl.nil?

  /(amb|blu|brn|gry|grn|hzl|oth)/.match? ecl
end

def valid_pid(passport)
  check_number(passport, 'pid', 9, 000_000_000, 999_999_999)
end

if $PROGRAM_NAME == __FILE__
  puts solve_first
  puts solve_second
end
