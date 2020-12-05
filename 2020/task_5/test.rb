require 'test/unit'
require_relative './task'

class Task5 < Test::Unit::TestCase
  def test_get_place
    assert_equal([44, 5], get_place('FBFBBFFRLR'))
    assert_equal([44, 5], get_place('FBFBBFFRLR\n'))
  end

  def test_get_id
    assert_equal(357, get_id([44, 5]))
  end

end
