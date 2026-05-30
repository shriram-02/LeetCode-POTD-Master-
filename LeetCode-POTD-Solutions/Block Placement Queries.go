const MAXN = 50000

var result [2.4 * MAXN]bool
var gap = &BIT{}
var blocks = &HiBitset{}

func getResults(queries [][]int) []bool {
	size := min(MAXN, 3*len(queries))
	gap.New(size + 1)
	blocks.New()
	blocks.Set(0)
	for _, q := range queries {
		if len(q) == 2 {
			blocks.Set(q[1])
		}
	}
	var prev int
	blocks.Iter(func(x int) {
		gap.Update(x, x-prev)
		prev = x
	})
	res := len(result)
	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		if len(q) == 2 {
			x := q[1]
			blocks.Unset(x)
			if after := blocks.Next(x + 1); after != -1 {
				gap.Update(after, after-blocks.Prev(x))
			}
		} else {
			x, sz := q[1], q[2]
			res--
			result[res] = x-blocks.Prev(x) >= sz || gap.Query(x, sz)
		}
	}
	return result[res:]
}

const SHIFT = 6
const MASK = 1<<SHIFT - 1
const L2 = (MAXN + MASK) >> SHIFT
const L1 = (L2 + MASK) >> SHIFT

type BIT struct {
	bit [MAXN + 1]uint32
	n   int
}

func (b *BIT) New(n int) {
	clear(b.bit[:b.n])
	b.n = n
}

func (b *BIT) Query(i, sz int) bool {
	var v uint32
	size := uint32(sz)
	for ; i >= 0 && v < size; i = (i & (i + 1)) - 1 {
		v = max(v, b.bit[i])
	}
	return v >= size
}

func (b *BIT) Update(i, v int) {
	vv := uint32(v)
	for ; i < b.n; i = i | (i + 1) {
		b.bit[i] = max(b.bit[i], vv)
	}
}

type HiBitset struct {
	l2   [L2]uint64
	l1   [L1]uint64
	l0   uint64
	maxv int
}

func (h *HiBitset) New() {
	size := h.maxv >> SHIFT
	clear(h.l2[:size+1])
	clear(h.l1[:size>>SHIFT+1])
	h.l0 = 0
	h.maxv = 0
}

func (h *HiBitset) Set(v int) {
	h.maxv = max(h.maxv, v)
	idx2 := v >> SHIFT
	idx1 := idx2 >> SHIFT
	h.l2[idx2] |= 1 << (v & MASK)
	h.l1[idx1] |= 1 << (idx2 & MASK)
	h.l0 |= 1 << (idx1 & MASK)
}

func (h *HiBitset) Unset(v int) {
	idx2 := v >> SHIFT
	idx1 := idx2 >> SHIFT
	h.l2[idx2] &^= 1 << (v & MASK)
	if h.l2[idx2] == 0 {
		h.l1[idx1] &^= 1 << (idx2 & MASK)
		if h.l1[idx1] == 0 {
			h.l0 &^= 1 << (idx1 & MASK)
		}
	}
}

func (h *HiBitset) Next(v int) int {
	idx2 := v >> SHIFT
	idx1 := idx2 >> SHIFT

	if next := h.l2[idx2] & (^uint64(0) << (v & MASK)); next != 0 {
		return (v &^ MASK) | bits.TrailingZeros64(next)
	}

	if next := h.l1[idx1] & (^uint64(0) << ((idx2 & MASK) + 1)); next != 0 {
		next2 := (idx1 << SHIFT) | bits.TrailingZeros64(next)
		return (next2 << SHIFT) | bits.TrailingZeros64(h.l2[next2])
	}

	if next := h.l0 & (^uint64(0) << ((idx1 & MASK) + 1)); next != 0 {
		next1 := bits.TrailingZeros64(next)
		next2 := (next1 << SHIFT) | bits.TrailingZeros64(h.l1[next1])
		return (next2 << SHIFT) | bits.TrailingZeros64(h.l2[next2])
	}

	return -1
}

func (h *HiBitset) Prev(v int) int {
	idx2 := v >> SHIFT
	idx1 := idx2 >> SHIFT

	if prev := h.l2[idx2] & ((1 << (v & MASK)) - 1); prev != 0 {
		return (v &^ MASK) | (bits.Len64(prev) - 1)
	}

	if prev := h.l1[idx1] & ((1 << (idx2 & MASK)) - 1); prev != 0 {
		prev2 := (idx1 << SHIFT) | (bits.Len64(prev) - 1)
		return (prev2 << SHIFT) | (bits.Len64(h.l2[prev2]) - 1)
	}

	if prev := h.l0 & ((1 << idx1) - 1); prev != 0 {
		prev1 := bits.Len64(prev) - 1
		prev2 := (prev1 << SHIFT) | (bits.Len64(h.l1[prev1]) - 1)
		return (prev2 << SHIFT) | (bits.Len64(h.l2[prev2]) - 1)
	}

	return -1
}

func (h *HiBitset) Iter(f func(v int)) {
	l0 := h.l0
	for l0 != 0 {
		idx1 := uint64(bits.TrailingZeros64(l0))
		l1 := h.l1[idx1]
		for l1 != 0 {
			off2 := uint64(bits.TrailingZeros64(l1))
			idx2 := (idx1 << SHIFT) | off2
			l2 := h.l2[idx2]
			for l2 != 0 {
				offv := uint64(bits.TrailingZeros64(l2))
				v := (idx2 << SHIFT) | offv
				f(int(v))
				l2 &= l2 - 1
			}
			l1 &= l1 - 1
		}
		l0 &= l0 - 1
	}
}