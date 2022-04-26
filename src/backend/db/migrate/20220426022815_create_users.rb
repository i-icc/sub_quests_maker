class CreateUsers < ActiveRecord::Migration[6.1]
  def change
    create_table :users do |t|
      t.int :id
      t.string :uid
      t.string :nickname

      t.timestamps
    end
  end
end
