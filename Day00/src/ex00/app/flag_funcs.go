package app

import (
	"fmt"
	"math"
	"sort"
	"slices"
)

//Среднее арифметическое
func (d *myData) meanAlg(){
	if len(d.Data) == 0 {
		fmt.Println("No mean")
		return
	}

	var result float64
	for _, value := range d.Data{
		result += value 
	}
	average := result / float64(d.quantity)
	average = math.Round(average*100)/100
	fmt.Printf("Mean: %.2f\n", average)
}
//Серединное значение.
func (d *myData) medianAlg(){
	if len(d.Data) == 0 {
		fmt.Println("No median")
		return
	}

	prec := d.quantity % 2
	sort.Float64s(d.Data)
	if prec == 0{
		i := (d.quantity / 2) - 1
		res := d.Data[i] + ((d.Data[i+1] - d.Data[i]) / 2)
		res = math.Round(res*100)/100
		fmt.Printf("Median: %.2f\n", res)
	}else{
		step := ((d.quantity - 1) / 2)
		fmt.Printf("Median: %.2f\n", d.Data[step])
	}
}
//Значение с наибольшей частотой
func (d *myData) modeAlg(){
	if len(d.Data) == 0 {
		fmt.Println("No mode")
		return
	}

	modeFreq := make(map[float64]int)
	//считаем повторения каждого числа
	for _, val := range d.Data{
		modeFreq[val]++
	}
	maxfreq := 0
	var modes []float64
	//Поиск наибольшего повторения и запоминаем числа с одинаковой частотой повт.
	for val, freq := range modeFreq{
		if freq > maxfreq{
			maxfreq = freq
			modes = []float64{val}		
		}
		if freq == maxfreq{
			modes = append(modes, val)
		}
	}
	if maxfreq <= 1{
		fmt.Println("No mode")
		return
	}
	minVal := slices.Min(modes)
	fmt.Printf("Mode: %.2f\n", minVal)
}
//StandardDeviationCalculation Среднеквадратическое отклонение — наиболее распространённый показатель рассеивания значений случайной величины
//относительно её математического ожидания (аналога среднего арифметического с бесконечным числом исходов).
//Обычно означает квадратный корень из дисперсии случайной величины, но иногда может означать тот или иной вариант оценки этого значения.
func (d *myData) sdAlg(){
	if len(d.Data) == 0 {
		fmt.Println("No SD")
		return
	}
	var sd float64
	//Average
	var result float64
	for _, value := range d.Data{
		result += value 
	}
	average := result / float64(d.quantity)
	//SD
	for _, val := range d.Data{
		sd += math.Pow(val-average, 2)
	}
	res := math.Sqrt(sd/float64(d.quantity))
	res = math.Round(res*100)/100
	fmt.Printf("SD: %.2f\n", res)
}