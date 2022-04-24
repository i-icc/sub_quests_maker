# sub_quests_maker
目的なく散歩するのも楽しいけどミッションとかあるともっと楽しいよね
そんなやつ

# setup
- docker-compose build
- docker-compose run backend bundle exec rails new . --force --database=mysql
- docker-compose run frontend npx create-nuxt-app@v2.15.0 