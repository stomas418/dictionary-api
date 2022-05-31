import mysql.connector
from mysql.connector import Error
from decouple import config

def create_connection(host_name, user_name, user_password,db_name):
    connection = None
    try:
        connection = mysql.connector.connect(
            host=host_name,
            user=user_name,
            passwd=user_password,
            database=db_name
        )
        print("Connection to MySQL DB successful")
    except Error as e:
        print(f"The error '{e}' occurred")
    return connection

def execute_query(connection:mysql.connector.MySQLConnection, query:str, args=None):
    cursor = connection.cursor()
    print(cursor, connection, query, args)
    cursor.execute(query, args)
    connection.commit()

def save_words_to_database(connection):
    import json
    def get_meanings(meanings_obj):
        meanings = []
        for key in meanings_obj:
            meanings.append(meanings_obj[key][1]) 
        
        return ";;".join(meanings)
    
    ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for letter in ALPHABET:
        filename = f"D{letter}.json"
        with open(filename, "r") as file:
            dictionary_section = json.load(file)
            create_table_query = f"CREATE TABLE IF NOT EXISTS {letter}(word CHAR(45) PRIMARY KEY UNIQUE, meanings MEDIUMTEXT)"
            execute_query(connection, create_table_query)
            for word in dictionary_section:
                save_query = f"INSERT INTO {letter}(word, meanings)VALUES (%s, %s)"
                meanings = get_meanings(dictionary_section[word]["MEANINGS"])
                values = (word, meanings)
                execute_query(connection, save_query, values)

def main():
    ENV_HOST = config("host")
    ENV_USER = config("user")
    ENV_PASSWORD = config("password")
    ENV_DATABASE = config("database")
    connection = create_connection(ENV_HOST, ENV_USER, ENV_PASSWORD, ENV_DATABASE)
    
    save_words_to_database(connection)
 
if __name__ == "__main__":
    main()