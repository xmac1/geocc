const fs = require('fs');
let fileNameList = fs.readdirSync('./geodata','utf-8');
let countries = [];
let brace = {countries};
for(let fileName of fileNameList) {
    let file = fs.readFileSync('./geodata/' + fileName,'utf-8');
    file = JSON.parse(file);
    let features = file.features;
    for(let item of features) {
        const properties = item.properties;
        const name = properties.cca2;
        const geometry = item.geometry;
        if(!geometry){
            continue;
        }
        const type = geometry.type;
        const coordinates = geometry.coordinates;
        const coordinatesParseList = type ==='Polygon'?[coordinates]:coordinates;
        
        for(let coordinateParseItem of coordinatesParseList) {

            coordinateParseItem = coordinateParseItem[0];
            let bounds = {x:null,y:null,width:null,height:null};
            let xMin = Number.MAX_SAFE_INTEGER,xMax = Number.MIN_SAFE_INTEGER,yMin = Number.MAX_SAFE_INTEGER,yMax = Number.MIN_SAFE_INTEGER;
            let out = {name,geo:null};
            let filterCoordinateParseItem = [];
            let filterNum = 2;
            coordinateParseItem.forEach((v,i) => {
                if(i % filterNum === 0) {
                    filterCoordinateParseItem.push(v.map(a => a*1e5>>0));
                }
            })
            out.geo = filterCoordinateParseItem;
            for(let point of filterCoordinateParseItem) {
                let x = point[0];
                let y = point[1];
                if(xMin > x) {
                    xMin = x;
                }
                if(xMax < x) {
                    xMax = x;
                }
                if(yMin > y) {
                    yMin = y;
                }
                if(yMax < y) {
                    yMax = y;
                }
            }
            bounds.x = xMin;
            bounds.y = yMin;
            bounds.width = (xMax-xMin) >> 0;
            bounds.height = (yMax-yMin) >> 0;
            out.bounds = bounds;
            countries.push(out);
            console.log({name,type,length:filterCoordinateParseItem.length,bounds});
        }
    }
}
fs.writeFileSync('./countries.json',JSON.stringify(brace));
