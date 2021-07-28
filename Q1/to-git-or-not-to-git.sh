content=$(curl -s https://01.kood.tech/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"maximihajlov\"}}){id}}"}')
id=$(jq -r '.data .user[0] .id' <<<"$content")
echo $id