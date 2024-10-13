# DCEL

## Doubly Connected Edge List


### DCEL 

- Vertex
- Half-edge
- Face

---

### Vertex

- id [identifier] 
- incident_halfedges [list of halfedes connected to the vertext]


### Half-edge

- origin [vertext]
- destination [vertext]
- incident face [face to the left side]
- twin [twin half edge]

### face 

- bountry [list of half edges in anti clockwise]
