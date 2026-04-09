using SQLite;

namespace AvaloniaApplication7.Services.Entities;


[Table("Songs")]
public class Song
{
    [PrimaryKey, AutoIncrement, Indexed]
    [Column("Id")]
    public int SongId { get; set; }
    public string Name { get; set; }
    public int Duration { get; set; }
    [Indexed]
    public int SingerId {get; set;}

}