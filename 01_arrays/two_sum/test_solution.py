import importlib.util
from pathlib import Path

import pytest

path = Path(__file__).with_name("solution.py")
spec = importlib.util.spec_from_file_location("two_sum_solution", path)
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)

two_sum = module.two_sum


@pytest.mark.parametrize(
    "nums,target,expected",
    [
        ([2, 7, 11, 15], 9, [0, 1]),
        ([3, 2, 4], 6, [1, 2]),
        ([3, 3], 6, [0, 1]),
        ([1, 2, 3], 10, []),
        ([], 5, []),
    ],
)
def test_two_sum(nums, target, expected):
    assert two_sum(nums, target) == expected
