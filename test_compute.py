import pytest
import compute

def test_compute_no_input():
    threshold = 0.0
    limit = 0.0
    values = []
    vsum, deltas = compute.compute(threshold, limit, values)
    assert vsum == 0.0
    assert len(deltas) == len(values)

def test_compute_one_input():
    threshold = 0.0
    limit = 0.0
    values = [0.0]
    vsum, deltas = compute.compute(threshold, limit, values)
    assert vsum == 0.0
    assert len(deltas) == len(values)
    
def test_compute_first_sample():
    threshold = 1000.0
    limit = 1000000.0
    values = [19.0, 0.0, 1000, 1001.5, 20000, 25000000.0]
    expected = [0.0, 0.0, 0.0, 1.5, 19000.0, 980998.5]
    vsum, deltas = compute.compute(threshold, limit, values)
    assert vsum == 1000000.0
    assert len(deltas) == len(values)
    for i in range(len(deltas)):
        assert deltas[i] == expected[i]

def test_compute_second_sample():
    threshold = 0.0
    limit = 100.0
    values = [0.0, 10.0, 50.0, 50.0, 10.0, 20.0]
    expected = [0.0, 10.0, 50.0, 40.0, 0.0, 0.0]
    vsum, deltas = compute.compute(threshold, limit, values)
    assert vsum == 100.0
    assert len(deltas) == len(values)
    for i in range(len(deltas)):
        assert deltas[i] == expected[i]
