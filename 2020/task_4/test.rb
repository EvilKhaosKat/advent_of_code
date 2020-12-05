require 'test/unit'
require_relative './task'

class Task4 < Test::Unit::TestCase
  def test_byr
    assert_true valid_byr({ 'byr' => '1921' })
    assert_false valid_byr({ 'byr' => '1919' })
    assert_false valid_byr({ 'byr' => '1dsf' })
  end

  def test_hgt
    assert_true valid_hgt({ 'hgt' => '60in' })
    assert_true valid_hgt({ 'hgt' => '190cm' })

    assert_false valid_hgt({ 'hgt' => '9001cm' })
    assert_false valid_hgt({ 'hgt' => '190in' })
    assert_false valid_hgt({ 'hgt' => '190' })
  end

  def test_hcl
    assert_true valid_hcl({ 'hcl' => '#f23456' })
    assert_false valid_hcl({ 'hcl' => '#zzzz42' })
    assert_false valid_hcl({ 'meow' => '#f23456' })
  end

  def test_ecl
    assert_true valid_ecl({ 'ecl' => 'grn' })
    assert_false valid_ecl({ 'ecl' => 'zzz' })
  end

  def test_second_valid_passport
    assert_true valid_second('pid' => '087499704',
                             'hgt' => '74in',
                             'ecl' => 'grn',
                             'iyr' => '2012',
                             'eyr' => '2030',
                             'byr' => '1980',
                             'hcl' => '#623a2f')
  end
end
