require "./spec_helper"

describe Part2 do
  it "wrong triangle fails check" do
    is_possible_triangle([5, 10, 25]).should eq(false)
  end

  it "correct triangle passes check" do
    is_possible_triangle([5, 5, 5]).should eq(true)
  end
end
