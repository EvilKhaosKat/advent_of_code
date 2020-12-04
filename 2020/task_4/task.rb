def solve_first
  passport = {}

  valid_passports = 0
  File.open('input.txt', 'r').each_line do |line|
    if line.length == 1
      valid_passports = valid_passports.next if valid passport
      passport = {}
    else
      line.split(' ').each do |pair|
        data = pair.split(':')
        passport[data[0]] = data[1]
      end
    end
  end
  valid_passports = valid_passports.next if valid passport

  valid_passports
end

# @param [Dict] passport
def valid(passport)
  passport.key?('byr') \
    && passport.key?('iyr') \
    && passport.key?('eyr') \
    && passport.key?('hgt') \
    && passport.key?('hcl') \
    && passport.key?('ecl') \
    && passport.key?('pid')
end

puts solve_first
#246 is too low