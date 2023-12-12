package main

import (
    "fmt"
    "maps"
)

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m) //palavra reservada

    v1 := m["k1"]
    fmt.Println("v1:", v1) // repasse de valor

    v3 := m["k3"]
    fmt.Println("v3:", v3) // zerado pq nao foi atribuido

    fmt.Println("len:", len(m)) // len palavra reservada

    delete(m, "k2")
    fmt.Println("map:", m)

    clear(m)
    fmt.Println("map:", m) // zera o obj ou array nao aprendi ainda

    _, prs := m["k2"]
    fmt.Println("prs:", prs) // preciso aprendder o _ mas acredito que por o array estar limpo retornou false, isso mesmo, comentei e testei

    n := map[string]int{"foo": 1, "bar": 2} // cria com atribuição de objetos
    fmt.Println("map:", n)

    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}