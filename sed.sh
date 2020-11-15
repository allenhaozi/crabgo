set -x
for file in `find . -type f -name "*.go"`
do
	sed -i -e 's/gitlab\.4pd\.io\/pht3\/app-manager/github\.com\/allenhaozi\/crabgo/g' $file
done
