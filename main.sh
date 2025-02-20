#!/bin/bash

awk -F',' 'NR>1 { 
    name=$2; 
    gsub(/ /, "_", name);
    stats="";
    stats=$5 "," $6 "," $9 "," $8 "," $7 "," $10; # Sheets vs Design need to swap some columns
    system("./ComSSA-Charts --vals " stats);
    system("mv ./out.png tmp/" name ".png");
}' tmp/data.csv

