package mop

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"testing"

	"github.com/Arafatk/glot"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"github.com/swkings/optimization/util"
)

func TestDominate(t *testing.T) {
	a := []int{1, 10}
	b := []int{2, 8}
	c := []int{2, 11}
	res1 := Dominate(a, b, true)
	res2 := Dominate(b, a, true)
	res3 := Dominate(a, c, true)
	res4 := Dominate(c, b, true)
	res5 := Dominate(a, a, true)
	assert.Equal(t, DominateRes(0), res1)
	assert.Equal(t, DominateRes(0), res2)
	assert.Equal(t, DominateRes(1), res3)
	assert.Equal(t, DominateRes(-1), res4)
	assert.Equal(t, AEqualB, res5)

	d := []int{1}
	e := []int{1}
	res6 := Dominate(d, e)
	assert.Equal(t, AEqualB, res6)
	f := []int{-1}
	res7 := Dominate(e, f)
	assert.Equal(t, BDominateA, res7)
}

func TestNDSet(t *testing.T) {
	data1 := [][]int{
		{6, 9},
		{5, 8},
		{4, 7},
		{3, 8},
		{1, 10},
		{2, 9},
		{1, 9},
	}
	fmt.Printf("data1: %v\n", data1)
	res1 := NDSet(data1)
	fmt.Printf("NDSetIndex1: %v\n", res1)
	ndset := [][]int{}
	rest := [][]int{}
	for i, item := range data1 {
		if util.ElementInList(i, res1) {
			ndset = append(ndset, item)
		} else {
			rest = append(rest, item)
		}

	}
	fmt.Printf("NDSet: %v\n", ndset)
	fmt.Printf("rest: %v\n", rest)

	data2 := [][]int{
		{1, 17, 5},
		{2, 1, 7},
		{2, 1, 6},
		{2, 7, 17},
		{2, 11, 12},
		{2, 19, 11},
		{14, 4, 9},
		{14, 12, 10},
		{15, 3, 13},
		{16, 12, 1},
		{17, 7, 6},
	}
	fmt.Printf("data2: %v\n", data2)
	res2 := NDSet(data2)
	fmt.Printf("data2: %v\n", data2)
	fmt.Printf("res2: %v\n", res2)
	fmt.Printf("NDSet: \n")
	for _, item := range res2 {
		fmt.Printf("%v\n", data2[item])
	}

	p := plot.New()
	p.Title.Text = "NDSet"
	p.X.Label.Text = "obj#1"
	p.Y.Label.Text = "obj#2"
	ndsetPoints := make(plotter.XYs, len(ndset))
	for i := range ndset {
		ndsetPoints[i].X, ndsetPoints[i].Y = float64(ndset[i][0]), float64(ndset[i][1])
	}
	restPoints := make(plotter.XYs, len(rest))
	fmt.Printf("rest: %v\n", rest)
	for i := range rest {
		restPoints[i].X, restPoints[i].Y = float64(rest[i][0]), float64(rest[i][1])
	}
	err := plotutil.AddScatters(p,
		"ndset", ndsetPoints,
		"rest", restPoints)
	if err != nil {
		log.Fatal(err)
	}

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "ndset.png"); err != nil {
		log.Fatal(err)
	}
}

func TestNDSetRank(t *testing.T) {
	data1 := [][]int{
		{6, 9},
		{5, 8},
		{4, 7},
		{3, 8},
		{1, 10},
		{2, 9},
		{1, 9},
	}
	fmt.Printf("data1: %v\n", data1)
	res1 := NDSetRank(data1)
	fmt.Printf("NDSetIndex1: %v\n", res1)
	ndset1 := [][][]int{}

	for _, ndL := range res1 {
		nd := [][]int{}
		for _, item := range ndL {
			nd = append(nd, data1[item])
		}
		ndset1 = append(ndset1, nd)
	}

	fmt.Printf("NDSet1: %v\n", ndset1)

	p := plot.New()
	p.Title.Text = "NDSet Rank"
	p.X.Label.Text = "obj#1"
	p.Y.Label.Text = "obj#2"

	var pointMap map[string]plotter.XYs = make(map[string]plotter.XYs)
	for rank, ndL := range ndset1 {
		ndsetPoints := make(plotter.XYs, len(ndL))
		for i, item := range ndL {
			ndsetPoints[i].X, ndsetPoints[i].Y = float64(item[0]), float64(item[1])
		}
		pointMap["rank-"+strconv.Itoa(rank+1)] = ndsetPoints
	}

	err := plotutil.AddScatters(p, util.MapUnpack(pointMap)...)
	if err != nil {
		log.Fatal(err)
	}

	if err = p.Save(8*vg.Inch, 8*vg.Inch, "./fig/ndset.rank.2d.png"); err != nil {
		log.Fatal(err)
	}

	data2 := [][]int{
		{1, 17, 5},
		{2, 1, 7},
		{2, 1, 6},
		{2, 7, 17},
		{2, 11, 12},
		{2, 19, 11},
		{14, 4, 9},
		{14, 12, 10},
		{15, 3, 13},
		{16, 12, 1},
		{17, 7, 6},
	}
	fmt.Printf("data2: %v\n", data2)
	res2 := NDSetRank(data2)
	fmt.Printf("data2: %v\n", data2)
	fmt.Printf("NDSetIndex2: %v\n", res2)
	ndset2 := [][][]int{}
	for _, ndL := range res2 {
		nd := [][]int{}
		for _, item := range ndL {
			nd = append(nd, data2[item])
		}
		ndset2 = append(ndset2, nd)
	}
	fmt.Printf("NDSet2: %v\n", ndset2)

	dimensions := 3
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	pointGroupName := "rank"
	style := "points"

	for rank, ndset := range ndset2 {
		points := make([][]int, len(ndset[0]))
		for _, item := range ndset {
			for di := 0; di < len(item); di++ {
				points[di] = append(points[di], item[di])
			}
		}
		name := pointGroupName + strconv.Itoa(rank)
		plot.AddPointGroup(name, style, points)
	}
	plot.SetTitle("NDSet Rank")
	plot.SetXLabel("obj#1")
	plot.SetYLabel("obj#2")
	plot.SetZLabel("obj#3")
	plot.SavePlot("./fig/ndset.rank.3d.png")
}

func TestSlice(t *testing.T) {
	a := []int{1}
	b := 1
	fmt.Printf("type: %v, %v\n", reflect.TypeOf(a), reflect.SliceOf(reflect.TypeOf(b)))
}

func TestHV(t *testing.T) {

	data2d := [][]int{
		{6, 9},
		{5, 8},
		{4, 7},
		{3, 8},
		{1, 10},
		{2, 9},
		{1, 9},
	}
	idealP2 := IdealPoint(data2d)
	fmt.Printf("idealP: %v\n", idealP2)
	nadirP2 := NadirPoint(data2d)
	fmt.Printf("nadirP: %v\n", nadirP2)
	refP2 := RefPoint(data2d)
	fmt.Printf("refP: %v\n", refP2)
	hv2 := HyperVolume2D(data2d)
	fmt.Printf("hv: %v\n", hv2)

	data3d := [][]int{
		{1, 17, 5},
		{1, 16, 6},
		{2, 1, 7},
		{2, 1, 6},
		{2, 7, 17},
		{2, 11, 12},
		{2, 19, 11},
		{14, 4, 9},
		{14, 12, 10},
		{15, 3, 13},
		{16, 12, 1},
		{17, 7, 6},
	}
	ps := NewParetoSolutionSet(data3d)
	idealP3 := ps.IdealPoint()
	fmt.Printf("idealP: %v\n", idealP3)
	nadirP3 := ps.NadirPoint()
	fmt.Printf("nadirP: %v\n", nadirP3)
	refP3 := ps.RefPoint()
	fmt.Printf("refP: %v\n", refP3)
	hv3 := ps.HyperVolume()
	fmt.Printf("hv: %v\n", hv3)
}
