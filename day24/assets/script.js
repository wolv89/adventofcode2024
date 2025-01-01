
const _OPS = {
    '38': 'AND',
    '94': 'XOR',
    '124': 'OR',
}

ruleset = document.getElementById('rules')

for (let r in _RULES) {

    let rule = _RULES[r]

    let inp1 = document.getElementById('gate-' + rule.inp1)
    let inp2 = document.getElementById('gate-' + rule.inp2)
    let out = document.getElementById('gate-' + rule.out)

    let inp1rect = inp1.getBoundingClientRect()
    let inp2rect = inp2.getBoundingClientRect()
    let outrect = out.getBoundingClientRect()

    let inp1x = Math.floor(inp1rect.x + (inp1rect.width / 2))
    let inp2x = Math.floor(inp2rect.x + (inp2rect.width / 2))
    let outx = Math.floor(outrect.x + (outrect.width / 2))

    let inp1y = Math.floor(inp1rect.y + inp1rect.height)
    let inp2y = Math.floor(inp2rect.y + inp2rect.height)
    let outy = Math.floor(outrect.y)

    let min_x = Math.min(inp1x, inp2x, outx)
    let max_x = Math.max(inp1x, inp2x, outx)
    let min_y = Math.min(inp1y, inp2y, outy)
    let max_y = Math.max(inp1y, inp2y, outy)

    // Hashtag DuaLipa
    let newrule = document.createElement('article')
    newrule.className = 'rule'
    newrule.id = `${rule.inp1}-${rule.inp2}-to-${rule.out}`

    newrule.style.left = min_x + 'px'
    newrule.style.top = min_y + 'px'
    newrule.style.width = max_x - min_x + 'px'
    newrule.style.height = max_y - min_y + 'px'

    let conntop = document.createElement('span')
    let connbottom = document.createElement('span')

    conntop.className = 'connection-top'
    connbottom.className = 'connection-bottom'

    let conntype = document.createElement('span')
    conntype.innerText = _OPS[rule.op]
    conntype.className = 'connection-type'

    // Connection Top Left/Width
    const ctl = Math.max(Math.min(inp1x, inp2x) - min_x, 0)
    const ctw = Math.abs(inp1x - inp2x)

    conntop.style.left = ctl + 'px'
    conntop.style.width = ctw + 'px'
    conntop.style.height = Math.floor(max_y - min_y) / 2 + 'px'

    const midp = Math.floor(ctl + ctw / 2)

    if (outx > midp) {
        connbottom.style.left = midp + 'px'
        connbottom.style.width = outx - midp - min_x + 'px'
        connbottom.className += ' right-side'
    } else {
        connbottom.style.left = '0px'
        connbottom.style.width = midp + 'px'
        connbottom.className += ' left-side'
    }

    connbottom.style.height = conntop.style.height

    conntop.append(conntype)
    newrule.append(conntop)
    newrule.append(connbottom)

    ruleset.append(newrule)

}
