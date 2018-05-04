function mklabel() {
	local exid=$1
	local unixtime=$(date +%s)

	printf "lola-%s-%s" "${exid}" "${unixtime}"
}
