KEY_NUMBER=$(cat streets/Buckingham_Place | sed -n '179p' | cut -d "#" -f 2)
echo $KEY_NUMBER
cat interviews/interview-$KEY_NUMBER
grep -A 5 "L337" vehicles | grep -A 4 -B 1 "Honda"| grep -A 2 -B 3 "Blue" | grep -B 4 "6'"
cat memberships/AAA memberships/Delta_SkyMiles memberships/Terminal_City_Library memberships/Museum_of_Bash_History | grep "$MAIN_SUSPECT" | wc -l