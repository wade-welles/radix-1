package radix

import "bytes"

const tabSize = 4

type edge struct {
	label string
	n     *Node
}

func (e *edge) writeTo(bd *builder, tabList []bool) {
	length := len(tabList)
	isLast, tlist := tabList[length-1], tabList[:length-1]
	for _, hasTab := range tlist {
		if hasTab {
			bd.Write(bytes.Repeat([]byte(" "), tabSize))
			continue
		}
		bd.WriteRune('│')
		bd.Write(bytes.Repeat([]byte(" "), tabSize-1))
	}
	if !isLast {
		bd.WriteRune('├')
	} else {
		bd.WriteRune('└')
	}
	bd.WriteString("── ")
	if bd.debug {
		bd.colors[colorRed].Fprintf(bd, "%d↑ ", e.n.priority)
	}
	bd.colors[colorBold].Fprintf(bd, "%s", e.label)
	if bd.debug {
		if e.n.IsLeaf() {
			bd.colors[colorGreen].Fprint(bd, " 🍂")
		}
		bd.colors[colorMagenta].Fprintf(bd, " → %#v", e.n.Value)
	}
	bd.WriteByte('\n')
	for i, next := range e.n.edges {
		if len(tabList) < next.n.depth { // runs only for the first edge
			tabList = append(tabList, i == len(e.n.edges)-1)
		} else {
			tabList[next.n.depth-1] = i == len(e.n.edges)-1
		}
		next.writeTo(bd, tabList)
	}
}
