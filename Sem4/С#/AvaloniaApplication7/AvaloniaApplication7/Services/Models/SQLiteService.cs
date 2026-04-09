using System;
using System.Collections.Generic;
using AvaloniaApplication7.Services.Contracts;
using AvaloniaApplication7.Services.Entities;
using RandomFriendlyNameGenerator;
using SQLite;

namespace AvaloniaApplication7.Services.Models;

public class SqLiteService :IDbService
{
    
    
    private SQLiteConnection? _db = null;
    
    
    public IEnumerable<Singer> GetAllSingers()
    {
        return _db.Table<Singer>().ToList();
    }

    public IEnumerable<Song> GetSingerSongs(int singerId)
    {
        
        var query = _db.Table<Song>().Where(s => s.SingerId == singerId);
        return query.ToList();
    }

    public void Init()
    {
        if (_db != null)
        {
            return;
        }

        _db = new SQLiteConnection("Data Source=AvaloniaApplication7.db");
        _db.CreateTable<Singer>();
        _db.CreateTable<Song>();


        var count = _db.Table<Singer>().Count();

        if (count == 0)
        {
            var groups = new[]
            {
                new Singer()
                {
                    Name = "Billy",
                    Country = "Brazil",
                    Id = 1,
                },
                new Singer()
                {
                    Name = "Frank",
                    Country = "France",
                    Id = 2,
                },
                new Singer()
                {
                    Name = "James",
                    Country = "Japan",
                    Id = 3,
                }
            };


            foreach (var g in groups)
            {
                _db.Insert(g);
            }


            Random rnd = new Random();

            for (int i = 0; i < groups.Length; i++)
            {
                for (int j = 0; j < rnd.Next(1,8); j++)
                {
                    var song = new Song()
                    {
                        Name = NameGenerator.PersonNames.Get(),
                        Duration = rnd.Next(1, 10),
                        SingerId = i + 1,
                        SongId = j + 1,
                    };

                    _db.Insert(song);
                }
            }
        }
    }
}