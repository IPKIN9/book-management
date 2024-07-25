FROM postgres:13

# Copy the script to create multiple databases
COPY init-db.sh /docker-entrypoint-initdb.d/
COPY custom-entrypoint.sh /usr/local/bin/

# Make the init-db.sh and custom-entrypoint.sh executable
RUN chmod +x /docker-entrypoint-initdb.d/init-db.sh /usr/local/bin/custom-entrypoint.sh

# Set the custom entrypoint
ENTRYPOINT ["custom-entrypoint.sh"]

# Use the default command for PostgreSQL
CMD ["postgres"]
