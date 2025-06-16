package visualization

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"

	"github.com/Soln1shko/heap-coursework/internal/heap"
	"github.com/fogleman/gg"
)

const (
	imageWidth  = 1200
	imageHeight = 800

	nodeRadius = 30

	backgroundColor = "#FFFFFF"
	nodeColor       = "#4CAF50"
	textColor       = "#FFFFFF"
	lineColor       = "#2196F3"
	highlightColor  = "#FF5722"
)

type HeapVisualizer struct {
	heap          *heap.Heap
	ctx           *gg.Context
	nodePositions []image.Point
}

func NewHeapVisualizer(h *heap.Heap) *HeapVisualizer {
	return &HeapVisualizer{
		heap:          h,
		ctx:           gg.NewContext(imageWidth, imageHeight),
		nodePositions: make([]image.Point, h.GetSize()),
	}
}

func (v *HeapVisualizer) DrawHeap() {
	v.ctx.SetHexColor(backgroundColor)
	v.ctx.Clear()

	maxLevel := 0
	size := v.heap.GetSize()
	for i := 0; (1<<i)-1 < size; i++ {
		maxLevel = i
	}

	v.drawEdges(0, 0, maxLevel)

	v.drawNodes()

	heapType := "Минимальная куча"
	if v.heap.Type == heap.MaxHeap {
		heapType = "Максимальная куча"
	}
	v.ctx.SetHexColor("#000000")
	v.ctx.LoadFontFace("Arial", 24)
	v.ctx.DrawStringAnchored(heapType, float64(imageWidth)/2, 30, 0.5, 0.5)
}

func (v *HeapVisualizer) drawNodes() {
	size := v.heap.GetSize()
	items := v.heap.GetItems()

	maxLevel := 0
	for i := 0; (1<<i)-1 < size; i++ {
		maxLevel = i
	}

	nodeIndex := 0
	for level := 0; level <= maxLevel && nodeIndex < size; level++ {
		nodesInLevel := 1 << level

		levelWidth := float64(imageWidth) * 0.8

		leftMargin := (float64(imageWidth) - levelWidth) / 2

		nodeSpacing := levelWidth / float64(nodesInLevel)

		y := 100 + float64(level)*120

		for i := 0; i < nodesInLevel && nodeIndex < size; i++ {
			x := leftMargin + nodeSpacing/2 + float64(i)*nodeSpacing

			v.nodePositions[nodeIndex] = image.Point{int(x), int(y)}

			v.ctx.SetHexColor(nodeColor)
			v.ctx.DrawCircle(x, y, nodeRadius)
			v.ctx.Fill()

			v.ctx.SetHexColor(textColor)
			v.ctx.LoadFontFace("Arial", 16)
			v.ctx.DrawStringAnchored(strconv.Itoa(items[nodeIndex]), x, y, 0.5, 0.5)

			nodeIndex++
		}
	}
}

func (v *HeapVisualizer) drawEdges(nodeIndex, level, maxLevel int) {
	size := v.heap.GetSize()
	if nodeIndex >= size {
		return
	}

	leftChildIndex := 2*nodeIndex + 1
	rightChildIndex := 2*nodeIndex + 2

	levelWidth := float64(imageWidth) * 0.8

	leftMargin := (float64(imageWidth) - levelWidth) / 2

	nodesInCurrentLevel := 1 << level

	nodesInNextLevel := 1 << (level + 1)

	currentNodeSpacing := levelWidth / float64(nodesInCurrentLevel)

	nextNodeSpacing := levelWidth / float64(nodesInNextLevel)

	nodeX := leftMargin + currentNodeSpacing/2 + float64(nodeIndex%(1<<level))*currentNodeSpacing
	nodeY := 100 + float64(level)*120

	if leftChildIndex < size {
		childX := leftMargin + nextNodeSpacing/2 + float64(leftChildIndex%(1<<(level+1)))*nextNodeSpacing
		childY := 100 + float64(level+1)*120

		v.ctx.SetHexColor(lineColor)
		v.ctx.SetLineWidth(2)
		v.ctx.DrawLine(nodeX, nodeY+nodeRadius, childX, childY-nodeRadius)
		v.ctx.Stroke()

		v.drawEdges(leftChildIndex, level+1, maxLevel)
	}

	if rightChildIndex < size {
		childX := leftMargin + nextNodeSpacing/2 + float64(rightChildIndex%(1<<(level+1)))*nextNodeSpacing
		childY := 100 + float64(level+1)*120

		v.ctx.SetHexColor(lineColor)
		v.ctx.SetLineWidth(2)
		v.ctx.DrawLine(nodeX, nodeY+nodeRadius, childX, childY-nodeRadius)
		v.ctx.Stroke()

		v.drawEdges(rightChildIndex, level+1, maxLevel)
	}
}

func (v *HeapVisualizer) SaveImage(filename string) error {
	return v.ctx.SavePNG(filename)
}
func (v *HeapVisualizer) HighlightNode(index int) {
	if index < 0 || index >= len(v.nodePositions) {
		return
	}

	pos := v.nodePositions[index]
	v.ctx.SetHexColor(highlightColor)
	v.ctx.DrawCircle(float64(pos.X), float64(pos.Y), nodeRadius+5)
	v.ctx.Stroke()
}

func CreateHeapImage(h *heap.Heap, filename string) error {
	visualizer := NewHeapVisualizer(h)
	visualizer.DrawHeap()
	return visualizer.SaveImage(filename)
}

func CreateHeapAnimationFrames(arr []int, heapType heap.HeapType, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	size := len(arr)
	h := heap.NewHeap(heapType, size)

	for i := 0; i < size; i++ {
		h.Insert(arr[i])

		filename := fmt.Sprintf("%s/frame_%03d.png", outputDir, i)
		if err := CreateHeapImage(h, filename); err != nil {
			return err
		}
	}

	return nil
}

func GenerateHeapVisualization(h *heap.Heap, outputPath string) error {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			img.Set(x, y, color.White)
		}
	}

	visualizer := NewHeapVisualizer(h)
	visualizer.DrawHeap()

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
